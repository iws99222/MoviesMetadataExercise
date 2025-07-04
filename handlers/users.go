package handlers

import (
	"encoding/json"
	"net/http" // importing net http
)

// this is a user's response struct
// it represent a user in the system
type allUsersResponse struct {
	// data of type User, named 'data' in the JSON response
	Data []User `json:"data"`
}

// handling all users endpoint
// this function will handle the /users endpoint
func HandleAllUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // set the response status code to 200 OK
	// setting json header
	w.Header().Set("Content-Type", "application/json")

	// now writing the response
	// those user will be
	response := allUsersResponse{
		// adding some preloaded users to the response
		Data: []User{
			{UserID: 1, Name: "Issa"},
			{UserID: 2, Name: "Ahmad"},
			{UserID: 3, Name: "Khaled"},
			{UserID: 4, Name: "Omar"},
		},
	}

	// encoding the response to JSON and writing it to the response writer
	json.NewEncoder(w).Encode(response)

}
