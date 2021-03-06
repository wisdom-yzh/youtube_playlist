package parser

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

const (
	METHOD    = "GET"
	BASE_URL  = "https://www.youtube.com/playlist"
	QUERY_KEY = "list"
	MATCH_REG = "var ytInitialData = (?P<InitialData>.*);</script>"
)

type Playlist struct {
	re *regexp.Regexp
}

func NewPlaylist() *Playlist {
	re, _ := regexp.Compile(MATCH_REG)
	return &Playlist{re: re}
}

func (client *Playlist) GetData(playlist string) (*YoutubePlayListData, error) {
	raw, err := client.getRawResponse(playlist)
	if err != nil {
		return nil, err
	}

	matchList := client.re.FindSubmatch(raw)
	initialDataIndex := client.re.SubexpIndex("InitialData")
	if initialDataIndex == -1 || len(matchList) <= initialDataIndex {
		return nil, errors.New("youtube data not found")
	}
	initialData := matchList[initialDataIndex]

	youtubePlayListData := &YoutubePlayListData{}
	if err := json.Unmarshal(initialData, youtubePlayListData); err != nil {
		return nil, err
	}
	return youtubePlayListData, nil
}

func (client *Playlist) getRawResponse(playlist string) ([]byte, error) {
	baseUrl, err := url.Parse(BASE_URL)
	if err != nil {
		return nil, err
	}

	query := baseUrl.Query()
	query.Add(QUERY_KEY, playlist)
	baseUrl.RawQuery = query.Encode()

	req, err := http.NewRequest(METHOD, baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	c := &http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}
