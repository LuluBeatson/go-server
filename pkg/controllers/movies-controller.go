package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/LuluBeatson/go-server/pkg/models"
	"github.com/gorilla/mux"
)

var movies []models.Movie

func init() {
	movies = append(movies, models.Movie{
		ID:    "1",
		Isbn:  "438227",
		Title: "Movie One",
		Director: &models.Director{
			Firstname: "John",
			Lastname:  "Doe",
		},
	})
	movies = append(movies, models.Movie{
		ID:    "2",
		Isbn:  "454555",
		Title: "Movie Two",
		Director: &models.Director{
			Firstname: "Steve",
			Lastname:  "Smith",
		},
	})
	fmt.Println("Movies initialized")
}

// get all movies as json
func GetMoviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// get a single movie by id
func GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	if movie, err := getMovie(params["id"]); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else {
		json.NewEncoder(w).Encode(*movie)
	}
}

// create a new movie
func CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie

	// decode the request body into a movie struct
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// generate a random id for the movie
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	createMovie(movie)

	json.NewEncoder(w).Encode(movie)
}

func UpdateMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var movie models.Movie

	// decode the request body into a movie struct
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := deleteMovie(params["id"]); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	movie.ID = params["id"]
	createMovie(movie)
	// TODO: make sure movies is unmodified in the event of an error

	json.NewEncoder(w).Encode(movie)
}

func DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if err := deleteMovie(params["id"]); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func getMovie(id string) (*models.Movie, error) {
	for _, movie := range movies {
		if movie.ID == id {
			return &movie, nil
		}
	}
	return nil, fmt.Errorf("Movie with id %v not found", id)
}

func createMovie(movie models.Movie) {
	movies = append(movies, movie)
}

func deleteMovie(id string) error {
	for i, movie := range movies {
		if movie.ID == id {
			// Delete movie from slice by appending all movies before and after the index
			movies = append(movies[:i], movies[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Movie with id %s not found", id)
}
