package main

import (
	"encoding/json"
	"fmt"
	"log"
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

var movies []Movie //var movie is slice of type movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies { //index = position of item in slice, item = actual k:v pair
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // deletes movie with id
			break
		}
	}
	json.NewEncoder(w).Encode(movies) //returns all movies after delete
}

func getSingleMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies { // '_' is iterator allows us to iterate without defining variable for iterator
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)     // decodes request body
	movie.ID = strconv.Itoa(rand.Intn(1000000000)) // generates integer value and converts it into string
	movies = append(movies, movie)                 // the latter movie points to the movie received in request
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // deletes movie with id
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie) // decodes request body
			movie.ID = params["id"]
			movies = append(movies, movie) // the latter movie points to the movie received in request
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "123456", Title: "RaOne", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "654321", Title: "OneRa", Director: &Director{Firstname: "Jane", Lastname: "Doe"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getSingleMovie).Methods("GET")
	r.HandleFunc("/addmovies", createMovie).Methods("POST")
	r.HandleFunc("/modmovies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/delmovies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
