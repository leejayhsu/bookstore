package handlers

import (
	"books/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetBook is the get handler
func GetBook(w http.ResponseWriter, r *http.Request) {
	bid := mux.Vars(r)["book_id"]
	i, err := strconv.Atoi(bid)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "invalid book_id"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")

	b := models.GetBook(i)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(b)
}
