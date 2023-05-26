package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Deepbaran/movie-rating-service/configs"
	"github.com/Deepbaran/movie-rating-service/routes"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// New DB Connection
	configs.ConnectDB()

	// API v1
	http.HandleFunc("/", routes.MovieV1)

	log.Fatal(http.ListenAndServe(port(), nil))
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
