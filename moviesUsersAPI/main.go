package main

import (
	"Moviestask/moviesUsersAPI/handlers"
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
			// here i want to check if i am getting all movies or a specific movie by id
			// will check if the query parameter 'id' is present, so we can fetch that movie by id
			movieID := r.URL.Query().Get("id")
			// getting the userID from the query url
			userID := r.URL.Query().Get("user_id")
			// checking if both userID and movieID are set
			if userID != "" && movieID != "" {
				// both userID and movieID are set, so we will call GetWatchedMoviesByUserID
				handlers.GetWatchedMoviesByUserID(w, r)
				return
			}

			// here i am checking if the movieID is set or not
			if movieID != "" {
				// the movieId is set, so will call the GetMovieByID func
				handlers.GetMovieByID(w, r)
				return // exiting since we already handled the request
			}

			// if the method is GET, we will call get all movies function
			handlers.GetAllMovies(w, r)

			// if the method is POST, we will add a new movie
		case http.MethodPost:
			handlers.AddMovie(w, r)

		case http.MethodDelete:
			// if the method is DELETE, we will delete a movie by ID
			handlers.DeleteMovie(w, r)
		default:
			// if the method is not GET or POST, we will return a 405 Method Not Allowed error
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		}
	})

	fmt.Println("Server is running on port 7777...")
	// starting the server on port 7777
	http.ListenAndServe(":7777", nil)
}
