// this function will be used to show the error in JSON format
package utils

import (
	"net/http"
)

// THIS IS REUSABLE FUNCTION

// i will use this function to show the error in JSON format
func ShowJsonError(w http.ResponseWriter, statusCode int, message string) {
	// setting json header
	w.Header().Set("Content-Type", "application/json")
	// setting the status code
	w.WriteHeader(statusCode)
	// now we will write the error message to the response
	w.Write([]byte(`{"error":"` + message + `"}`)) // writing the error message in JSON format

}
