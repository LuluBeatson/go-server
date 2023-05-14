package routes

import (
	"github.com/LuluBeatson/go-server/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterMoviesRoutes = func(router *mux.Router) {
	router.HandleFunc("/movies", controllers.GetMoviesHandler).Methods("GET")
	router.HandleFunc("/movies/{id}", controllers.GetMovieHandler).Methods("GET")
	router.HandleFunc("/movies", controllers.CreateMovieHandler).Methods("POST")
	router.HandleFunc("/movies/{id}", controllers.UpdateMovieHandler).Methods("PUT")
	router.HandleFunc("/movies/{id}", controllers.DeleteMovieHandler).Methods("DELETE")
}
