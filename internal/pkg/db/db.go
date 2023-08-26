package db

import (
	"os"

	"github.com/airelcamilo/podvoyage-backend/internal/pkg/utils"
	"github.com/airelcamilo/podvoyage-backend/internal/podvoyage/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := os.Getenv("SUPABASE_DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	utils.CheckErrIsNil(err)

	db.AutoMigrate(&model.Category{}, &model.Folder{}, &model.Podcast{}, &model.Episode{}, &model.Item{}, &model.Queue{})
	podcast := model.Podcast{
		Id: 1,
	}
	if result := db.FirstOrCreate(&podcast); result.Error != nil {
		return nil
	}

	if result := db.Where("podcast_id = ?", 1).Delete(&model.Item{}); result.Error != nil {
		return nil
	}

	return db
}
