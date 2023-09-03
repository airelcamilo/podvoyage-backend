package podvoyageService

import (
	"errors"

	"github.com/airelcamilo/podvoyage-backend/internal/podvoyage/model"
	"gorm.io/gorm"
)

type QueueService struct {
	DB *gorm.DB
}

func (s *QueueService) New(db *gorm.DB) QueueService {
	return QueueService{db}
}

func (s *QueueService) GetAllQueue(userId int) ([]model.Queue, error) {
	var queue []model.Queue
	if result := s.DB.Preload("Episode").Where("user_id = ?", userId).Find(&queue); result.Error != nil {
		return queue, result.Error
	}
	return queue, nil
}

func (s *QueueService) AddToQueue(request *model.Episode, userId int) (model.Queue, error) {
	episode := *request
	var temp model.Queue
	var queue model.Queue
	if episode.Id == 0 {
		episode.PodcastId = 1
		episode.UserId = 0
		if result := s.DB.Create(&episode); result.Error != nil {
			return queue, result.Error
		}
	}

	if result := s.DB.Where("episode_id = ? AND user_id = ?", episode.Id, userId).First(&temp); result.Error != nil {
		queue = model.Queue{
			Episode:   &episode,
			UserId:    userId,
			EpisodeId: episode.Id,
		}
		if result := s.DB.Create(&queue); result.Error != nil {
			return queue, result.Error
		}
		return queue, nil
	}
	return queue, errors.New("podcast already queue")
}

func (s *QueueService) RemoveInQueue(id int, userId int) error {
	var queue model.Queue
	if result := s.DB.Where("user_id = ?", userId).First(&queue, id).Delete(&queue); result.Error != nil {
		return result.Error
	}
	return nil
}