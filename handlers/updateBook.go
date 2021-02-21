package handlers

import (
	"books/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// UpdateBook is a PUT function
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bid := mux.Vars(r)["book_id"]
	id, err := strconv.Atoi(bid)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "invalid book_id"}`))
		return
	}
	var b models.Book
	err = json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b.ID = id
	models.UpdateBook(b)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "You updated a book!"}`))
}

// PatchBook is a PATCH handler
func PatchBook(w http.ResponseWriter, r *http.Request) {
	bid := mux.Vars(r)["book_id"]
	id, err := strconv.Atoi(bid)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "invalid book_id"}`))
		return
	}

	var payload map[string]interface{}
	d := json.NewDecoder(r.Body)
	d.UseNumber()
	if err := d.Decode(&payload); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request"}`))
		return
	}
	book := models.PatchBook(id, payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
