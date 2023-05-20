package main

import (
	"database/sql"
	"log"

	gin "github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"

	routes "github.com/Deepbaran/movie-rating-service/routes"
)

var DB *sql.DB

var logger = log.Default()

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./movies.db")
	// defer db.Close()
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func main() {
	// Create the Database connection
	if err := ConnectDatabase(); err != nil {
		logger.Fatal(err)
	}

	router := gin.Default()

	// API v1
	v1 := router.Group("/api/v1")
	{
		v1.GET("/movies", routes.GetMovies)
		v1.GET("/movies/:id", routes.GetMovie)
		v1.POST("/movies", routes.CreateMovie)
		v1.PATCH("/movies/:id", routes.UpdateMovie)
		v1.DELETE("/movies/:id", routes.DeleteMovie)
	}

	// It will serve on PORT 8080
	// PORT environment variable is set as 8080 (The default value)
	router.Run(":8080")
}
