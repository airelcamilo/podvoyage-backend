package model

import "gorm.io/gorm"

type Folder struct {
	Id         int        `json:"id" gorm:"primaryKey"`
	FolderName string     `json:"folderName"`
	Podcasts   []*Podcast `json:"podcasts" gorm:"many2many:folders_podcasts;constraint:OnDelete:CASCADE;"`
}

func (f *Folder) AfterCreate(tx *gorm.DB) (err error) {
	if result := tx.Model(&Item{}).Create(&Item{
		Type:       "Folder",
		Name:       f.FolderName,
		ArtistName: "You",
		ArtworkUrl: "",
		FolderId:   f.Id,
	}); result.Error != nil {
		return result.Error
	}
	return nil
}

func (f *Folder) AfterDelete(tx *gorm.DB) (err error) {
	var item Item
	if result := tx.Model(&Item{}).Where("folder_id = ?", f.Id).First(&item); result.Error != nil {
		return result.Error
	}

	if result := tx.Model(&Item{}).Delete(&item); result.Error != nil {
		return result.Error
	}
	return nil
}

type FolderPodcast struct {
	FolderId  int `json:"podcastId" gorm:"primaryKey"`
	PodcastId int `json:"folderId" gorm:"primaryKey"`
}
