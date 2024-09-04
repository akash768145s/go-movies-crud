package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies  []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index,item:=range movies {
		if item.ID==params["id"] {
			movies=append(movies[:index],movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _,item:=range movies {
		if item.ID==params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_=json.NewDecoder(r.Body).Decode(&movie)
	movie.ID=strconv.Itoa(rand.Intn(100000000))
	movies=append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index,item:=range movies {
		if item.ID==params["id"] {
			movies=append(movies[:index],movies[index+1:]...)
			var movie Movie
			_=json.NewDecoder(r.Body).Decode(&movie)
			movie.ID=params["id"]
			movies=append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return 
		}
	}
}

func main() {
	r:=mux.NewRouter()

	movies=append(movies, Movie{ID:"1", Isbn:"1234567890", Title:"The Shawshank Redemption", Director:&Director{Firstname:"Frank", Lastname:"Darabont"}})
	movies=append(movies, Movie{ID:"2", Isbn:"1234567891", Title:"The Godfather", Director:&Director{Firstname:"Francis", Lastname:"Ford"}})
	
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting Server at Port 8000\n")
	log.Fatal(http.ListenAndServe(":8080", r))

}