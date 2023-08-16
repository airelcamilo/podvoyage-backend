package podvoyageService

import (
	"github.com/airelcamilo/podvoyage-backend/internal/podvoyage/model"
	"gorm.io/gorm"
)

type EpisodeService struct {
	DB *gorm.DB
}

func (s *EpisodeService) New(db *gorm.DB) EpisodeService {
	return EpisodeService{db}
}

func (s *EpisodeService) MarkAsPlayed(id int) (model.Episode, error) {
	var episode model.Episode
	if result := s.DB.First(&episode, id).Update("played", !(episode.Played)); result.Error != nil {
		return episode, result.Error
	}
	return episode, nil
}

func (s *EpisodeService) SetCurrentTime(id int, request *model.CurrentTimeRequest) (model.Episode, error) {
	var episode model.Episode
	if result := s.DB.First(&episode, id).Update("current_time", request.CurrentTime); result.Error != nil {
		return episode, result.Error
	}
	return episode, nil
}