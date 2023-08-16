package model

type Item struct {
	Id         int    `json:"id" gorm:"primaryKey"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	ArtistName string `json:"artistName"`
	ArtworkUrl string `json:"artworkUrl"`
	PodcastId  int    `json:"podcastId"`
	TrackId    int    `json:"trackId"`
	FolderId   int    `json:"folderId"`
	Pos        int    `json:"pos" gorm:"autoIncrement:1;autoIncrementIncrement:1"`
}
