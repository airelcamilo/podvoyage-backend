package podvoyageService

import (
	"github.com/airelcamilo/podvoyage-backend/internal/podvoyage/model"
	"gorm.io/gorm"
)

type ItemService struct {
	DB *gorm.DB
}

func (s *ItemService) New(db *gorm.DB) ItemService {
	return ItemService{db}
}

func (s *ItemService) GetAllItem() ([]model.Item, error) {
	var items []model.Item
	if result := s.DB.Find(&items); result.Error != nil {
		return items, result.Error
	}
	return items, nil
}

// inc / dec pos

