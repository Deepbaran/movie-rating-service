package configs

import (
	"os"

	"github.com/Deepbaran/movie-rating-service/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_NAME")), &gorm.Config{})
	if err != nil {
		panic("Failed to connect with database")
	}

	err = db.AutoMigrate((&models.Movie{}))
	if err != nil {
		return
	}

	DB = db
}
