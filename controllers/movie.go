package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Deepbaran/movie-rating-service/configs"
	"github.com/Deepbaran/movie-rating-service/models"
	"github.com/gin-gonic/gin"
)

// Schema for creating movies
type CreateMovieInput struct {
	MovieId     string `json:"movie_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	Rating      int    `json:"rating"`
}

// Schema for updating movies
type UpdateMovieInput struct {
	Rating int `json:"rating"`
}

func CreateMovie(c *gin.Context) {
	// Validate input
	var input CreateMovieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Movie
	movie := models.Movie{
		MovieId:     strings.Join(strings.Split(input.Name, " "), ""),
		Name:        input.Name,
		Description: input.Description,
		Genre:       input.Genre,
		Rating:      input.Rating,
	}

	if err := configs.DB.Where("name=?", movie.Name).Find(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := configs.DB.Create(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record already exists!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

func GetMovies(c *gin.Context) {
	var movies []models.Movie
	if err := configs.DB.First(&movies).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No records found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": movies})
}

func GetMovie(c *gin.Context) {
	movie_id := c.Param("movie_id")
	fmt.Println(movie_id)
	var movie models.Movie
	if err := configs.DB.Where("movie_id=?", movie_id).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

func UpdateMovie(c *gin.Context) {
	movie_id := c.Param("movie_id")
	var movie models.Movie
	if err := configs.DB.Where("movie_id=?", movie_id).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateMovieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := configs.DB.Model(&movie).Update("rating", input.Rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not created!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

func DeleteMovie(c *gin.Context) {
	movie_id := c.Param("movie_id")
	var movie models.Movie
	if err := configs.DB.Where("movie_id=?", movie_id).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := configs.DB.Delete(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not deleted!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": movie})
}
