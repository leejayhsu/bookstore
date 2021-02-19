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
	i, _ := strconv.Atoi(bid)
	w.Header().Set("Content-Type", "application/json")

	b := models.GetBook(i)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(b)
}
