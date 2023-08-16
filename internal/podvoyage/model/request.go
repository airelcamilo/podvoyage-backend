package model

type SearchRequest struct {
	PodcastName string `json:"podcastName"`
}

type CurrentTimeRequest struct {
	CurrentTime string `json:"currentTime"`
}
