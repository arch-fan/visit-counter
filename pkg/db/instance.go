package db

import (
	"os"
	"visit-counter/pkg/db/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(os.Getenv("SQLITE_FILE_PATH")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Page{})

	return db, nil
}
