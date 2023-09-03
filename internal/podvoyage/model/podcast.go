package model

import (
	"gorm.io/gorm"
)

type Podcast struct {
	Id          int         `json:"id" gorm:"primaryKey"`
	TrackId     int         `json:"trackId"`
	UserId      int         `json:"userId"`
	PodcastName string      `json:"trackName"`
	ArtistName  string      `json:"artistName"`
	FeedUrl     string      `json:"feedUrl"`
	ArtworkUrl  string      `json:"artworkUrl600"`
	Desc        string      `json:"desc" xml:"channel>description"`
	Website     string      `json:"link" xml:"channel>link"`
	Categories  []*Category `json:"categories" xml:"channel>category" gorm:"many2many:podcasts_categories;constraint:OnDelete:CASCADE;"`
	Episodes    []*Episode  `json:"episodes" xml:"channel>item" gorm:"foreignKey:PodcastId;references:Id;constraint:OnDelete:CASCADE;"`
}

type Category struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" xml:"text,attr"`
}

func (p *Podcast) AfterCreate(tx *gorm.DB) (err error) {
	if result := tx.Model(&Item{}).Create(&Item{
		UserId:     p.UserId,
		Type:       "Podcast",
		Name:       p.PodcastName,
		ArtistName: p.ArtistName,
		ArtworkUrl: p.ArtworkUrl,
		PodcastId:  p.Id,
		TrackId:    p.TrackId,
	}); result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *Podcast) AfterDelete(tx *gorm.DB) (err error) {
	var item Item
	if result := tx.Model(&Item{}).Where("podcast_id = ? AND user_id = ?", p.Id, p.UserId).First(&item); result.Error != nil {
		return nil
	}

	if result := tx.Model(&Item{}).Delete(&item); result.Error != nil {
		return result.Error
	}
	return nil
}
