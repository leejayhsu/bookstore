package main

import (
	"books/handlers"
	"books/models"
	"fmt"
	"log"
	"net/http"
	"os"

	"database/sql"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	// Initalize the sql.DB connection pool and assign it to the models.DB
	// global variable.
	var err error
	models.DB, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer models.DB.Close()

	// Create the "books" table.
	if _, err := models.DB.Exec(dbinit); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting the bookstore")
	r := mux.NewRouter()
	v1 := r.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/books/{book_id}", handlers.GetBook).Methods(http.MethodGet)
	v1.HandleFunc("/books", handlers.CreateBook).Methods(http.MethodPost)
	v1.HandleFunc("/books/{book_id}", handlers.UpdateBook).Methods(http.MethodPut)
	v1.HandleFunc("/books/{book_id}", handlers.PatchBook).Methods(http.MethodPatch)
	v1.HandleFunc("/books/{book_id}", handlers.DeleteBook).Methods(http.MethodDelete)
	v1.HandleFunc("/", handlers.HomeLink)

	log.Fatal(http.ListenAndServe("localhost:5000", r))
}

var dbinit string = `
CREATE TABLE IF NOT EXISTS
books (
	id SERIAL PRIMARY KEY,
	title varchar(255),
	price numeric(20,2)
)
`
