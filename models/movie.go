package models

type Movie struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	Rating      int    `json:'rating"`
}
