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
	id := models.InsertBook(b)

	response := struct {
		ID    int
		Title string
		Price float32
	}{
		id, b.Title, b.Price,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
