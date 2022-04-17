package parser

type YoutubePlayListData struct {
	Contents struct {
		TwoColumnBrowseResultsRenderer TwoColumnBrowseResultsRenderer `json:"twoColumnBrowseResultsRenderer"`
	} `json:"contents"`
	Metadata Metadata `json:"metadata"`
}
type Thumbnails struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
type Thumbnail struct {
	Thumbnails []Thumbnails `json:"thumbnails"`
}
type Runs struct {
	Text string `json:"text"`
}
type AccessibilityData struct {
	Label string `json:"label"`
}
type Accessibility struct {
	AccessibilityData AccessibilityData `json:"accessibilityData"`
}
type Title struct {
	Runs          []Runs        `json:"runs"`
	Accessibility Accessibility `json:"accessibility"`
}
type Index struct {
	SimpleText string `json:"simpleText"`
}
type ShortBylineText struct {
	Runs []Runs `json:"runs"`
}
type LengthText struct {
	Accessibility Accessibility `json:"accessibility"`
	SimpleText    string        `json:"simpleText"`
}
type WebCommandMetadata struct {
	URL         string `json:"url"`
	WebPageType string `json:"webPageType"`
	RootVe      int    `json:"rootVe"`
}
type CommandMetadata struct {
	WebCommandMetadata WebCommandMetadata `json:"webCommandMetadata"`
}
type VssLoggingContext struct {
	SerializedContextData string `json:"serializedContextData"`
}
type LoggingContext struct {
	VssLoggingContext VssLoggingContext `json:"vssLoggingContext"`
}
type CommonConfig struct {
	URL string `json:"url"`
}
type HTML5PlaybackOnesieConfig struct {
	CommonConfig CommonConfig `json:"commonConfig"`
}
type WatchEndpointSupportedOnesieConfig struct {
	HTML5PlaybackOnesieConfig HTML5PlaybackOnesieConfig `json:"html5PlaybackOnesieConfig"`
}
type WatchEndpoint struct {
	VideoID                            string                             `json:"videoId"`
	PlaylistID                         string                             `json:"playlistId"`
	Index                              int                                `json:"index"`
	Params                             string                             `json:"params"`
	LoggingContext                     LoggingContext                     `json:"loggingContext"`
	WatchEndpointSupportedOnesieConfig WatchEndpointSupportedOnesieConfig `json:"watchEndpointSupportedOnesieConfig"`
}
type NavigationEndpoint struct {
	ClickTrackingParams string          `json:"clickTrackingParams"`
	CommandMetadata     CommandMetadata `json:"commandMetadata"`
	WatchEndpoint       WatchEndpoint   `json:"watchEndpoint"`
}
type PlaylistVideoRenderer struct {
	VideoID            string             `json:"videoId"`
	Thumbnail          Thumbnail          `json:"thumbnail"`
	Title              Title              `json:"title"`
	Index              Index              `json:"index"`
	ShortBylineText    ShortBylineText    `json:"shortBylineText"`
	LengthText         LengthText         `json:"lengthText"`
	NavigationEndpoint NavigationEndpoint `json:"navigationEndpoint"`
	LengthSeconds      string             `json:"lengthSeconds"`
	TrackingParams     string             `json:"trackingParams"`
	IsPlayable         bool               `json:"isPlayable"`
}
type PlaylistVideoListRenderer struct {
	Contents []struct {
		PlaylistVideoRenderer PlaylistVideoRenderer `json:"playlistVideoRenderer"`
	} `json:"contents"`
}
type ItemSectionRenderer struct {
	Contents []struct {
		PlaylistVideoListRenderer PlaylistVideoListRenderer `json:"playlistVideoListRenderer"`
	} `json:"contents"`
}
type SectionListRenderer struct {
	Contents []struct {
		ItemSectionRenderer ItemSectionRenderer `json:"itemSectionRenderer"`
	} `json:"contents"`
}
type Content struct {
	SectionListRenderer SectionListRenderer `json:"sectionListRenderer"`
}
type TabRenderer struct {
	Selected bool    `json:"selected"`
	Content  Content `json:"content"`
}
type Tabs struct {
	TabRenderer TabRenderer `json:"tabRenderer"`
}
type TwoColumnBrowseResultsRenderer struct {
	Tabs []Tabs `json:"tabs"`
}
type PlaylistMetadataRenderer struct {
	Title string `json:"title"`
}
type Metadata struct {
	PlaylistMetadataRenderer PlaylistMetadataRenderer `json:"playlistMetadataRenderer"`
}
