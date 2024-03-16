package channel

import "google.golang.org/api/youtube/v3"

type Channel struct {
	Id           string
	Title        string
	ChannelId    string
	ThumbnailUrl string
	Views        uint64
}

func Format(i int, response *youtube.ChannelListResponse) Channel {
	return Channel{
		Id:           response.Items[i].Id,
		Title:        response.Items[i].Snippet.Title,
		ChannelId:    response.Items[i].Id,
		ThumbnailUrl: response.Items[i].Snippet.Thumbnails.Default.Url,
		Views:        response.Items[i].Statistics.ViewCount,
	}
}
