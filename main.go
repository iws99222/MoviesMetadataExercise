package main

import (
	"Moviestask/handlers"
	"Moviestask/utils"
	"fmt"
	"net/http"
)

func main() {
	// Register the handler for the /Users endpoint

	// Preloaded User data
	http.HandleFunc("/users", handlers.HandleAllUsers)

	// getting all movies
	http.HandleFunc("/movies", func(w http.ResponseWriter, r *http.Request) {
		// here i am checking the method of the request
		// since we are using the same endpoint for both GET and POST requests
		switch r.Method {
		case http.MethodGet:
			// this is the handler which will be controlling wether i want to get all movies
			// or  get a specific movie by ID
			handlers.HandleGetMovies(w, r)
		case http.MethodPost:
			handlers.AddMovie(w, r)
		case http.MethodDelete:
			// if the method is DELETE, we will delete a movie by ID
			handlers.DeleteMovie(w, r)
		default:
			// IF THE METHOD IS NOT GET OR POST OR DELETE, WE WILL RETURN THE JSON ERROR
			utils.ShowJsonError(w, http.StatusMethodNotAllowed, "Method is not allowed. Please use either GET, POST. or DELETE.")

		}
	})
	// handling getting a specific movie by ID
	http.HandleFunc("/movie", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetMovieByID(w, r)
	})

	fmt.Println("Server is running on port 7777...")
	// starting the server on port 7777
	http.ListenAndServe(":7777", nil) // Start the server on port 7777
}
