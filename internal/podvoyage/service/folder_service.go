package podvoyageService

import (
	"github.com/airelcamilo/podvoyage-backend/internal/podvoyage/model"
	"gorm.io/gorm"
)

type FolderService struct {
	DB *gorm.DB
}

func (s *FolderService) New(db *gorm.DB) FolderService {
	return FolderService{db}
}

func (s *FolderService) GetAllFolder(userId int) ([]model.Folder, error) {
	var folders []model.Folder
	if result := s.DB.Where("user_id = ?", userId).Find(&folders); result.Error != nil {
		return folders, result.Error
	}
	return folders, nil
}

func (s *FolderService) GetFolder(id int, userId int) (model.Folder, error) {
	var folder model.Folder
	if result := s.DB.Preload("Podcasts").Where("user_id = ?", userId).First(&folder, id); result.Error != nil {
		return folder, result.Error
	}
	return folder, nil
}

func (s *FolderService) SaveFolder(request *model.Folder, userId int) (model.Folder, error) {
	folder := *request
	folder.UserId = userId
	if result := s.DB.Create(&folder); result.Error != nil {
		return folder, result.Error
	}
	return folder, nil
}

func (s *FolderService) CheckInFolder(id int, userId int) (int, error) {
	var folder model.FolderPodcast
	var podcast model.Podcast

	if result := s.DB.Where("id = ? AND user_id = ?", id, userId).First(&podcast); result.Error != nil {
		return -1, result.Error
	}

	if result := s.DB.Table("folders_podcasts").Where("podcast_id = ?", id).First(&folder); result.Error != nil {
		return 0, nil
	}
	return folder.FolderId, nil
}

func (s *FolderService) ChangeFolder(folderId int, podId int, userId int) (int, error) {
	var folder model.Folder
	var oldFolder model.Folder
	var podcast model.Podcast
	var item1 model.Item
	var item2 model.Item

	if result := s.DB.Where("id = ? AND user_id = ?", podId, userId).First(&podcast); result.Error != nil {
		return -1, result.Error
	}

	oldFolderId, err := s.CheckInFolder(podId, userId)

	if oldFolderId == folderId {
		return folderId, nil
	}

	if oldFolderId == -1 {
		return -1, err
	} else if oldFolderId == 0 {
		if result := s.DB.Model(&item1).Where("podcast_id = ? AND user_id = ?", podId, userId).First(&item1).Delete(&item1); result.Error != nil {
			return -1, result.Error
		}
	} else {
		if result := s.DB.Where("id = ? AND user_id = ?", oldFolderId, userId).First(&oldFolder); result.Error != nil {
			return -1, result.Error
		}

		if err := s.DB.Model(&oldFolder).Association("Podcasts").Delete(&podcast); err != nil {
			return -1, err
		}
	}

	if folderId == 0 {
		if result := s.DB.Create(&model.Item{
			UserId:     userId,
			Type:       "Podcast",
			Name:       podcast.PodcastName,
			ArtworkUrl: podcast.ArtworkUrl,
			PodcastId:  podcast.Id,
			ArtistName: podcast.ArtistName,
		}); result.Error != nil {
			return -1, result.Error
		}
	} else {
		if result := s.DB.Where("id = ? AND user_id = ?", folderId, userId).First(&folder); result.Error != nil {
			return -1, result.Error
		}

		if err := s.DB.Model(&folder).Association("Podcasts").Append(&podcast); err != nil {
			return -1, err
		}

		if result := s.DB.Model(&item2).Where("podcast_id = ? AND user_id = ?", podId, userId).First(&item2).Delete(&item2); result.Error != nil {
			return -1, result.Error
		}
	}
	return folderId, nil
}

func (s *FolderService) RemoveFolder(id int, userId int) error {
	var folder model.Folder
	if result := s.DB.Where("user_id = ?", userId).First(&folder, id).Delete(&folder); result.Error != nil {
		return result.Error
	}
	return nil
}
