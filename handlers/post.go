package handlers

import (
	"books/models"
	"encoding/json"
	"net/http"
)

// CreateBook is the post handler
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var b models.Book
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	book := models.InsertBook(b)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}
