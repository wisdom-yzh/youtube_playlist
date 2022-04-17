package parser

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type VideoData struct {
	Thumbnails      []Thumbnails `json:"thumbnails"`
	VideoID         string       `json:"vid"`
	Name            string       `json:"name"`
	Label           string       `json:"label"`
	DurationSeconds int          `json:"durationSeconds"`
	Duration        string       `json:"duration"`
	DurationLabel   string       `json:"durationLabel"`
}

type PlayListData struct {
	Title string      `json:"title"`
	List  []VideoData `json:"list"`
}

func PlaylistHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	playlist, exist := params["list"]
	if !exist {
		http.Error(w, "params illegal", http.StatusBadRequest)
		return
	}

	rawData, err := NewPlaylist(playlist).GetData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(rawData.Contents.TwoColumnBrowseResultsRenderer.Tabs) == 0 {
		http.Error(w, "No data found in playlist", http.StatusBadRequest)
		return
	}

	contents := rawData.Contents.TwoColumnBrowseResultsRenderer.Tabs[0].TabRenderer.Content.SectionListRenderer.Contents[0].ItemSectionRenderer.Contents[0].PlaylistVideoListRenderer.Contents
	list := make([]VideoData, len(contents))

	for idx, content := range contents {
		inner := &content.PlaylistVideoRenderer
		list[idx].Thumbnails = inner.Thumbnail.Thumbnails
		list[idx].VideoID = inner.VideoID
		list[idx].Name = inner.Title.Runs[0].Text
		list[idx].Label = inner.Title.Accessibility.AccessibilityData.Label
		list[idx].DurationSeconds, _ = strconv.Atoi(inner.LengthSeconds)
		list[idx].Duration = inner.LengthText.SimpleText
		list[idx].DurationLabel = inner.LengthText.Accessibility.AccessibilityData.Label
	}

	data := &PlayListData{Title: rawData.Metadata.PlaylistMetadataRenderer.Title, List: list}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func VideoUrlHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	videoId, exist := params["video"]
	if !exist {
		http.Error(w, "params illegal", http.StatusBadRequest)
		return
	}

	url, err := GetDownloadUrl(videoId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(map[string]string{"url": url}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
