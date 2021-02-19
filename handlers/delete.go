package handlers

import "net/http"

// DeleteBook is the post handler
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "you deleted a book!"}`))
}
