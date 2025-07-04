package handlers

import (
	"Moviestask/utils"
	"encoding/json" // importing encoding json
	"net/http"

	"github.com/google/uuid"
)

// all messages response

var (
	movies      = make(map[string]Movie)           // map to store movies
	watchedList = make(map[string]map[string]bool) // map to store watched movies by user ID
)

// function to add a new movie
func AddMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	// decoding the request body to the movie struct
	// if there is an error in decoding, return a bad request error
	// if the request body is not a valid JSON, it will return an error
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		// showing json error
		utils.ShowJsonError(w, http.StatusBadRequest, "Invalid request body")
		return //exit the function if decoding fails
	}
	// checking if the movie has a title and year
	if movie.Title == "" || movie.Year == 0 {
		utils.ShowJsonError(w, http.StatusBadRequest, "Title and Year are required")
		return // exit the function if validation fails
	}

	// init a new UUID for the movie
	// using github.com/google/uuid package to generate a unique ID
	movieID := uuid.NewString()
	//setting the ID of the movie
	movie.ID = movieID      // setting the ID of the movie to the generated UUID
	movies[movieID] = movie // adding the movie to the map with a unique ID
	// setting header to application/json
	w.Header().Set("Content-Type", "application/json")
	// setting the response status code to 201 Created
	w.WriteHeader(http.StatusCreated)

	// encoding the movies map to JSON and writing it to the response writer
	json.NewEncoder(w).Encode(movie)
}

// function to delete a movie by ID
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	movieID := r.URL.Query().Get("id") // getting the movie ID from the query parameters
	// so example: /movies?id="112"

	// checking if the movie ID is provided and not null or empty
	if movieID == "" {
		utils.ShowJsonError(w, http.StatusBadRequest, "Movie ID is required in order to delete a movie")
		return // exiting the function if the movie ID is not provided
	}
	// checking the movie ID exists in the map
	if _, exists := movies[movieID]; !exists {
		utils.ShowJsonError(w, http.StatusNotFound, "Movie not found") // returning a 404 Not Found error if the movie ID does not exist in the map
		return                                                         // exiting
	}
	// else here, the movie ID exists in the map, so we can delete it
	// setting the header to application/json
	w.Header().Set("Content-Type", "application/json")
	// deleting the movie from the map
	delete(movies, movieID)
	w.WriteHeader(http.StatusNoContent) // setting the response status code to 204 since not content to return
}

// function to get a movie by ID
func GetMovieByID(w http.ResponseWriter, r *http.Request) {
	// first getting the movie id from the query url
	movieID := r.URL.Query().Get("id")

	// checking if the movie id is parsed in the url
	if movieID == "" {
		utils.ShowJsonError(w, http.StatusBadRequest, "Missing movie ID") // returning a 400 bad request error
		return                                                            // exiting the function if the movie id is not provided
	}

	// secondly, i want to check if the movie id exists in the map
	if movie, exists := movies[movieID]; !exists {
		utils.ShowJsonError(w, http.StatusBadRequest, "Movie you're trying to get does not exist") // returning bad request error
		return
	} else {
		// else here, the movie exists in the map
		// so we can return it to the user
		// setting the header to application/json
		w.Header().Set("Content-Type", "application/json")
		// now we are going to write the response
		w.WriteHeader(http.StatusOK) // setting the response status code to 200 which means OK
		// now will encode the movie to JSON and write it to the response writer
		json.NewEncoder(w).Encode(movie)

	}
}

// function to handle getting all the movies
func HandleGetMovies(w http.ResponseWriter, r *http.Request) {
	// getting the movieID from the query url
	movieID := r.URL.Query().Get("id")
	// getting the userID from the query
	userID := r.URL.Query().Get("user_id")
	// checking if both userID and movieID are set
	switch {
	case userID != "" && movieID != "":
		// both userID and movieID are set, so we will call GetWatchedMoviesByUserID
		getWatchedMoviesByUserID(w, r)
		return // exiting since we already handled the request
	case movieID != "":
		// the movieID is set, so we will call GetMovieByID
		GetMovieByID(w, r)
		return // exiting since we already handled the request
	default:
		// by default, if no userID or movieID is set, we will call GetAllMovies
		getAllMovies(w)
	}
}

// I SET THOSE FUNCTIONS TO BE PRIVATE, by starting their names with a lowercase letter
// because they will be handled by the function handleGetMovies above
// function to get movie with user id to check if the user watched the movie or not
func getWatchedMoviesByUserID(w http.ResponseWriter, r *http.Request) {
	// first getting the user id from the query url
	userID := r.URL.Query().Get("user_id")
	// secondly, getting the moivie id from the query url
	movieID := r.URL.Query().Get("id")
	// checking if the user id is parsed in the url
	if userID == "" || movieID == "" {
		utils.ShowJsonError(w, http.StatusBadRequest, "Missing user ID or Movie ID")
		return // exiting the function if the user id or is not provided
	}

	// checking if the user id and movie id exists in the watchedList
	if watchedList[userID][movieID] {
		// setting the header to application/json
		w.Header().Set("Content-Type", "application/json")
		// user has watched the movie
		// watched : true
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]bool{"watched": true})
	} else {
		// if the user has not watched the movie, we will return the movie details
		GetMovieByID(w, r)
		// and we will mark the movie as watched by the user
		go func() {
			if watchedList[userID] == nil {
				watchedList[userID] = make(map[string]bool)
			}
			watchedList[userID][movieID] = true // marking the movie as watched by the user
		}()
	}

}

// function to get all movies
func getAllMovies(w http.ResponseWriter) {
	// adding json header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// getting all movies, encoding them to json, and writing to the response writer
	json.NewEncoder(w).Encode(movies)
}
