package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Rating string `json:"rating"`
	Year   int    `json:"year"`
	Actor  *Actor `json:"actor"`
}

type Actor struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := &Response{Data: movies, Message: "Success", Code: 1000}
	json.NewEncoder(w).Encode(response)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(&Response{Data: item, Message: "Success", Code: 1000})
			return
		}
	}
	json.NewEncoder(w).Encode(&Response{Data: nil, Message: "Success", Code: 1001})

}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(&Response{Data: movie, Message: "Success", Code: 1000})
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:idx], movies[idx+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(&Response{Data: movie, Message: "Success", Code: 1000})
			return
		}
	}
	json.NewEncoder(w).Encode(&Response{Data: nil, Message: "Id not found", Code: 1002})
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:idx], movies[idx+1:]...)
			response := &Response{Data: movies, Message: "Success", Code: 1000}
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	json.NewEncoder(w).Encode(&Response{Data: nil, Message: "Id not found", Code: 1002})
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{Id: "1", Title: "Ayo pulang", Rating: "9", Year: 2019, Actor: &Actor{Name: "Rahmat", Age: 23}})

	// Route handles & endpoints
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	//start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
