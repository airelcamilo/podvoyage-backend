package db

import (
	"os"

	"github.com/airelcamilo/podvoyage-backend/internal/pkg/utils"
	pm "github.com/airelcamilo/podvoyage-backend/internal/podvoyage/model"
	um "github.com/airelcamilo/podvoyage-backend/internal/user/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db_addr := os.Getenv("SUPABASE_DB")
	db, err := gorm.Open(postgres.Open(db_addr), &gorm.Config{})
	utils.CheckErrIsNil(err)

	db.AutoMigrate(&pm.Category{}, &pm.Folder{}, &pm.Podcast{}, &pm.Episode{}, &pm.Item{}, &pm.Queue{})
	db.AutoMigrate(&um.User{}, &um.Session{})
	podcast := pm.Podcast{
		Id: 1,
	}
	if result := db.FirstOrCreate(&podcast); result.Error != nil {
		return nil
	}

	if result := db.Where("podcast_id = ?", 1).Delete(&pm.Item{}); result.Error != nil {
		return nil
	}

	return db
}
