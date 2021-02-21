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
	d := json.NewDecoder(r.Body)
	if err := d.Decode(&b); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request"}`))
		return
	}
	book := models.InsertBook(b)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}
