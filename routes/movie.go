package routes

import (
	"net/http"

	"github.com/Deepbaran/movie-rating-service/controllers"
)

func MovieV1(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// Get Query Parameter
	m := r.Form.Get("movie_id")

	switch r.URL.Path {
	case "/api/v1/movies":
		if r.Method == "GET" && len(m) > 0 {
			controllers.GetMovie(w, r, m)
		} else if r.Method == "GET" {
			controllers.GetMovies(w, r)
		} else if r.Method == "POST" {
			controllers.CreateMovie(w, r)
		} else if r.Method == "PUT" {
			controllers.UpdateMovie(w, r, m)
		} else if r.Method == "DELETE" {
			controllers.DeleteMovie(w, r, m)
		}
	default:
		http.Error(w, "Invalid Request.", http.StatusMethodNotAllowed)
	}
}
