package handlers

import "net/http"

// UpdateBook is the get handler
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "You updated a book!"}`))
}