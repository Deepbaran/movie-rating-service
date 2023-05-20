package models

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("./movies.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect with database")
	}

	err = db.AutoMigrate((&Movie{}))
	if err != nil {
		return
	}

	DB = db
}
