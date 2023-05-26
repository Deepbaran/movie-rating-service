package models

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	MovieId     string `json:"movie_id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	Rating      int    `json:"rating"`
}

func (movie *Movie) JsonMarshal() string {
	data, err := json.Marshal(*movie)
	if err != nil {
		fmt.Println("error: " + err.Error())
		return ""
	}
	return string(data)
}

// Insert

// Get All Movies

// Get a Specific Movie

// Update a Movie

// Delete a Movie
