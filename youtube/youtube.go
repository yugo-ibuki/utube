package youtube

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
	"os"
)

var (
	// チャンネルIDを指定
	ChannelIds = []string{""}
)

var listOptions = []string{"snippet", "contentDetails", "statistics"}

type Youtube struct {
	Service *youtube.Service
}

func New() *Youtube {
	apiKey := os.Getenv("YOUTUBE_API_KEY")

	s, err := youtube.NewService(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	return &Youtube{
		Service: s,
	}
}

func (y *Youtube) GetChCall(chId string) (*youtube.ChannelListResponse, error) {
	call := y.Service.Channels.List(listOptions).Id(chId)
	response, err := call.Do()
	if err != nil {
		return nil, err
	}
	return response, nil
}
