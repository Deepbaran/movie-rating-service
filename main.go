package main

import (
	"database/sql"

	gin "github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"

	routers "github.com/Deepbaran/movie-rating-service/routes"
)

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./movies.db")
	defer db.Close()
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func main() {
	// Create the Database connection
	if err := ConnectDatabase(); err != nil {
		panic(err)
	}

	router := gin.Default()

	// API v1
	v1 := router.Group("/api/v1")
	{
		v1.GET("/movies", routers.GetMovies)
		v1.GET("/movies/:id", routers.GetMovie)
		v1.POST("/movies", routers.CreateMovie)
		v1.PATCH("/movies/:id", routers.UpdateMovie)
		v1.DELETE("/movies/:id", routers.DeleteMovie)
	}

	// It will serve on PORT 8080
	// PORT environment variable is set as 8080 (The default value)
	router.Run(":8080")
}
