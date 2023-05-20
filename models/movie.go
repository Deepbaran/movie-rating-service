package models

type Movie struct {
	MovieId     string `json:"movie_id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	Rating      int    `json:"rating"`
}
