package handlers

import (
	"net/http"
)

// HomeLink is a dumb andler for the root endpoint
func HomeLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "welcome to the book store"}`))
}
