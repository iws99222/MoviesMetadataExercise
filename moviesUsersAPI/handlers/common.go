package handlers

// User struct represents a user in the system
type User struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}

// Message struct represents a message in the system
type Movie struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}
