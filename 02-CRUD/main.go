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

var movie []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" || r.URL.Path != "/movies" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	if len(movie) == 0 {
		fmt.Fprintf(w, "No Movies Found")
	} else {
		json.NewEncoder(w).Encode(movie)
	}

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" || r.URL.Path != "/movies/{id}" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movie {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" || r.URL.Path != "/movies" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var newMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&newMovie)
	newMovie.ID = strconv.Itoa(rand.Intn(1000000))
	movie = append(movie, newMovie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movie {
		if item.ID == params["id"] {
			movie = append(movie[:index], movie[index+1:]...)
			var newMovie Movie
			_ = json.NewDecoder(r.Body).Decode(&newMovie)
			newMovie.ID = params["id"]
			movie = append(movie, newMovie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	movieId := mux.Vars(r)["id"]
	for index, item := range movie {
		if item.ID == movieId {
			movie = append(movie[:index], movie[index+1:]...)
			break
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)

}

func main() {
	fmt.Println("Hello, World!")
	r := mux.NewRouter() // r is now a router object

	// Hardcoded data
	movie = append(movie, Movie{ID: "1", Isbn: "448743", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movie = append(movie, Movie{ID: "2", Isbn: "847564", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	// Route Handlers / Endpoints
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// Start the server
	fmt.Println("Server is running on port 8000")

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
}
