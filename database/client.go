package database

import (
	"vote-for-a-language/database/models"
	"vote-for-a-language/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Client *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open(utils.GetEnv("DATABASE_NAME")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Language{})

	Client = db
}
