package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func init() {
	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "438227",
		Title: "Movie One",
		Director: &Director{
			Firstname: "John",
			Lastname:  "Doe",
		},
	})
	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "454555",
		Title: "Movie Two",
		Director: &Director{
			Firstname: "Steve",
			Lastname:  "Smith",
		},
	})
	fmt.Println("Movies initialized")
}

// get all movies as json
func getMoviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// get a single movie by id
func getMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

// create a new movie
func createMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie

	// decode the request body into a movie struct
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		fmt.Println("Error decoding request body:", err)
	}
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	createMovie(movie)

	json.NewEncoder(w).Encode(movie)
}

func createMovie(movie Movie) {
	movies = append(movies, movie)
}

func updateMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var movie Movie

	// decode the request body into a movie struct
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		fmt.Println("Error decoding request body:", err)
	}

	deleteMovie(params["id"])
	movie.ID = params["id"]
	createMovie(movie)
	// TODO: make sure movies is unmodified in the event of an error

	json.NewEncoder(w).Encode(movie)
}

func deleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	deleteMovie(params["id"])
}

func deleteMovie(id string) {
	for i, movie := range movies {
		if movie.ID == id {
			// Delete movie from slice by appending all movies before and after the index
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
}
