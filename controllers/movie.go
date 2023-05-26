package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/Deepbaran/movie-rating-service/configs"
	"github.com/Deepbaran/movie-rating-service/models"
)

// Schema for creating movies
type CreateMovieInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	Rating      int    `json:"rating"`
}

// Schema for updating movies
type UpdateMovieInput struct {
	Description string `json:"description"`
	Genre       string `json:"genre"`
	Rating      int    `json:"rating"`
}

// PUT
// func CreateMovie(c *gin.Context) {
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	// Get the request body
	r.ParseForm()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	// Validate input
	var input CreateMovieInput
	if err := json.Unmarshal([]byte(body), &input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
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

	if err := configs.DB.Create(&movie).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Record already exists!")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, movie.JsonMarshal())
}

// GET
func GetMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie
	if err := configs.DB.First(&movies).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "No records found!")
		return
	}

	data := ""
	for i := 0; i < len(movies); i++ {
		data += movies[0].JsonMarshal()
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, data)
}

// GET
func GetMovie(w http.ResponseWriter, r *http.Request, m string) {
	var movie models.Movie
	if err := configs.DB.Where("movie_id=?", m).First(&movie).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Record not found!")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, movie.JsonMarshal())
}

// POST
func UpdateMovie(w http.ResponseWriter, r *http.Request, m string) {
	var movie models.Movie
	if err := configs.DB.Where("movie_id=?", m).First(&movie).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Record not found!")
		return
	}

	// Get the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	// Validate input
	var input UpdateMovieInput
	if err := json.Unmarshal([]byte(body), &input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	if len(input.Description) == 0 {
		input.Description = movie.Description
	}
	if len(input.Genre) == 0 {
		input.Genre = movie.Genre
	}

	if err := configs.DB.Model(&movie).Updates(input).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Record not created!")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, movie.JsonMarshal())
}

// DELETE
func DeleteMovie(w http.ResponseWriter, r *http.Request, m string) {
	var movie models.Movie
	if err := configs.DB.Where("movie_id=?", m).First(&movie).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Record not found!")
		return
	}

	if err := configs.DB.Delete(&movie).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Record not deleted!")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, movie.JsonMarshal())
}
