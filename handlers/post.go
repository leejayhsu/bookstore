package handlers

import (
	"books/models"
	"encoding/json"
	"fmt"
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

	w.WriteHeader(http.StatusCreated)

	response := []byte(fmt.Sprintf(`{"bookId":%d, "title":%q, "price":%f}`, id, b.Title, b.Price))
	w.Write(response)
}
