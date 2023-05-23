package main

import (
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Deepbaran/movie-rating-service/controllers"
	"github.com/Deepbaran/movie-rating-service/configs"
)

func main() {
	router := gin.Default()

	// New DB Connection
	configs.ConnectDB()

	// API v1
	v1 := router.Group("/api/v1")
	{
		v1.POST("/movies", controllers.CreateMovie)
		v1.GET("/movies", controllers.GetMovies)
		v1.GET("/movies/:movie_id", controllers.GetMovie)
		v1.PATCH("/movies/:movie_id", controllers.UpdateMovie)
		v1.DELETE("/movies/:movie_id", controllers.DeleteMovie)
	}

	// It will serve on PORT 8080
	// PORT environment variable is set as 8080 (The default value)
	router.Run(":8080")
}
