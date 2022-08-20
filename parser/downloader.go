package parser

import (
	"context"
	"errors"
	"fmt"
	"log"

	youtube "github.com/kkdai/youtube/v2"
)

func GetDownloadUrl(videoID string) (string, error) {
	ctx := context.Background()
	client := getClient()

	video, err := client.GetVideoContext(ctx, fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID))
	if err != nil {
		log.Printf("Failed to get video url from videoID %s", videoID)
		return "", err
	}

	audioFormats := video.Formats.Type("audio")

	var url string
	for _, format := range audioFormats {
		log.Printf("Check audio format: %v", &format)
		streamURL, err := client.GetStreamURLContext(ctx, video, &format)
		if streamURL == "" {
			log.Printf("Failed to parse audio format from video %s , error: %v", videoID, err)
			continue
		}
		url = streamURL
		break
	}

	if url == "" {
		errorMessage := fmt.Sprintf("Failed to get any audio formats from video %s", videoID)
		log.Println(errorMessage)
		return "", errors.New(errorMessage)
	}

	log.Printf("Get video url, videoID=%s, url=%s", videoID, url)
	return url, nil
}

func getClient() *youtube.Client {
	return &youtube.Client{}
}
