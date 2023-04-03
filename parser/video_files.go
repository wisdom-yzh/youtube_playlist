package parser

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"net/http"
)

type VideoFileStorage struct {
	VideoId string
	Headers http.Header
	Data    []byte
}

const pathPrefix = "/tmp/"

func SaveVideo(videoId string, headers http.Header, data []byte) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(VideoFileStorage{
		VideoId: videoId,
		Headers: headers,
		Data:    data,
	}); err != nil {
		log.Printf("Failed to encode video, err %v\n", err)
		return err
	}

	// Save the encoded binary data to a file
	if err := ioutil.WriteFile(pathPrefix+videoId, buf.Bytes(), 0644); err != nil {
		log.Printf("Failed to write video data to file, err %v\n", err)
		return err
	}

	log.Printf("Success to save video data to local file storage %s\n", pathPrefix+videoId)
	return nil
}

func LoadVideo(videoId string) (headers http.Header, videoData []byte, err error) {
	// Read the encoded binary data from the file
	data, err := ioutil.ReadFile(pathPrefix + videoId)
	if err != nil {
		log.Printf("Failed to read video data from videoId %s, err %v\n", videoId, err)
		return nil, nil, err
	}

	// Create a decoder
	dec := gob.NewDecoder(bytes.NewReader(data))

	// Decode the data
	var video VideoFileStorage
	err = dec.Decode(&video)
	if err != nil {
		log.Printf("Failed to unmarshal video data from videoId %s, err %v\n", videoId, err)
		return nil, nil, err
	}

	log.Println("Success to get video data from local file storage")
	return video.Headers, video.Data, nil
}
