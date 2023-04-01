package parser

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
)

var (
	videoManager = NewVideoStreamManager()
)

type VideoStreamManager struct {
	mu          sync.Mutex
	videoMap    map[string]*VideoStream
	recycleChan chan string
}

type VideoStream struct {
	mu          sync.Mutex
	videoId     string // key
	headers     http.Header
	body        io.ReadCloser
	data        []byte
	receivers   []*VideoStreamReceiver
	recycleChan chan string
}

type VideoStreamReceiver struct {
	headers  http.Header
	dataCh   chan []byte
	position int // position index to sent video stream data
	closed   atomic.Bool
}

func NewVideoStreamManager() *VideoStreamManager {
	vsm := &VideoStreamManager{
		mu:          sync.Mutex{},
		videoMap:    map[string]*VideoStream{},
		recycleChan: make(chan string),
	}

	go func(vsm *VideoStreamManager) {
		for {
			videoId := <-vsm.recycleChan
			vsm.removeVideoStream(videoId)
		}
	}(vsm)

	return vsm
}

func (vsm *VideoStreamManager) GetVideoStream(videoId string, url string) (*VideoStreamReceiver, error) {
	vsm.mu.Lock()
	if vs, ok := vsm.videoMap[videoId]; ok {
		vsm.mu.Unlock()
		log.Printf("Get exist video stream %s object from manger\n", videoId)
		return vs.AddReceiver(), nil
	}

	vs, err := vsm.createVideoStream(videoId, url)
	if err != nil {
		vsm.mu.Unlock()
		return nil, err
	}

	vsm.videoMap[videoId] = vs
	vsm.mu.Unlock()
	return vs.AddReceiver(), nil
}

func (vsm *VideoStreamManager) createVideoStream(videoId string, url string) (*VideoStream, error) {
	log.Printf("Create new video stream %s\n", videoId)

	res, err := http.Get(url)
	if err != nil {
		log.Printf("Create video stream failed: %v\n", err)
		return nil, err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusPartialContent {
		err = fmt.Errorf("Create video stream failed: Receive unexpected status code %d", res.StatusCode)
		log.Println(err.Error())
		return nil, err
	}

	return NewVideoStream(videoId, res.Header, res.Body, vsm.recycleChan), nil
}

func (vsm *VideoStreamManager) removeVideoStream(videoId string) {
	log.Printf("Remove video stream %s\n", videoId)
	vsm.mu.Lock()
	delete(vsm.videoMap, videoId)
	vsm.mu.Unlock()
}

func NewVideoStream(videoId string, headers http.Header, body io.ReadCloser, recycleChan chan string) *VideoStream {
	vs := &VideoStream{
		mu:          sync.Mutex{},
		videoId:     videoId,
		headers:     headers,
		data:        make([]byte, 0),
		body:        body,
		recycleChan: recycleChan,
		receivers:   make([]*VideoStreamReceiver, 0),
	}

	go vs.consumeVideoStream()
	return vs
}

func (vs *VideoStream) AddReceiver() *VideoStreamReceiver {
	vs.mu.Lock()
	defer vs.mu.Unlock()

	log.Printf("Add new receiver for video stream %s\n", vs.videoId)
	receiver := &VideoStreamReceiver{
		headers:  vs.headers,
		dataCh:   make(chan []byte, 1),
		position: len(vs.data),
		closed:   atomic.Bool{},
	}
	vs.receivers = append(vs.receivers, receiver)
	if receiver.position > 0 {
		receiver.dataCh <- vs.data[:len(vs.data)]
	}
	return receiver
}

type ConsumeStatus int

var (
	Consuming      ConsumeStatus = 0
	ConsumeSuccess ConsumeStatus = 1
	ConsumeError   ConsumeStatus = 2
)

func (vs *VideoStream) consumeVideoStream() {
	buf := make([]byte, 0, 1024*1024)
	var status ConsumeStatus

	// consume video stream chunks
	for status = Consuming; status == Consuming; {
		if len(buf) == cap(buf) {
			buf = append(buf, 0)[:len(buf)]
		}

		n, err := vs.body.Read(buf[len(buf):cap(buf)])
		log.Printf("Received chunk sized %d from %s", n, vs.videoId)
		buf = buf[:len(buf)+n]

		if err != nil {
			if err == io.EOF {
				// all video data consumed
				log.Println("All video data have been consumed")
				status = ConsumeStatus(ConsumeSuccess)
			} else {
				// error occurred
				log.Printf("Error occurred duration consume data: %v\n", err)
				status = ConsumeStatus(ConsumeError)
			}
		}

		vs.mu.Lock()
		vs.data = buf
		for _, receiver := range vs.receivers {
			if !receiver.IsClosed() {
				receiver.dataCh <- vs.data[receiver.position:len(buf)]
				receiver.position = len(buf)
			}
		}
		vs.mu.Unlock()
	}

	// if all data consumed normally, we will save video data and headers to local file,
	// and then close all receivers
	vs.mu.Lock()
	vs.body.Close()
	for _, receiver := range vs.receivers {
		close(receiver.dataCh)
	}
	if status == ConsumeSuccess {
		vs.saveHeadersAndData()
	}
	vs.mu.Unlock()

	// recycle
	vs.recycleChan <- vs.videoId
}

func (vs *VideoStream) saveHeadersAndData() {
	SaveVideo(vs.videoId, vs.headers, vs.data)
}

func (vsr *VideoStreamReceiver) GetReceiveChan() chan []byte {
	return vsr.dataCh
}

func (vsr *VideoStreamReceiver) GetHeaders() http.Header {
	return vsr.headers
}

func (vsr *VideoStreamReceiver) Close() {
	vsr.closed.Store(true)
}

func (vsr *VideoStreamReceiver) IsClosed() bool {
	return vsr.closed.Load()
}
