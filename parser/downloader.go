package parser

import (
	"log"
	"os"

	downloader "github.com/kkdai/youtube/v2/downloader"
)

func GetDownloadUrl(videoID string) (string, error) {
	client := getDownloader()
	video, err := getDownloader().GetVideo(videoID)
	if err != nil {
		return "", err
	}

	audioFormats := video.Formats.Type("audio")
	audioFormats.Sort()

	url, err := client.GetStreamURL(video, &audioFormats[0])
	if err != nil {
		return "", err
	}

	log.Println(videoID, "get video url", url)
	return url, nil
}

func getDownloader() *downloader.Downloader {
	return &downloader.Downloader{
		OutputDir: os.TempDir(),
	}
}
