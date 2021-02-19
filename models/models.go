package models

import (
	"database/sql"
	"fmt"
)

// DB Create an exported global variable to hold the database connection pool.
var DB *sql.DB

// Book is a model of a book
type Book struct {
	Title string
	Price float32
}

// InsertBook inserts a book
func InsertBook(b Book) int {
	sqlStatement := `
INSERT INTO books (title, price)
VALUES ($1, $2)
RETURNING id`
	id := 0
	err := DB.QueryRow(sqlStatement, b.Title, b.Price).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}

// GetBook retrieves a single book
func GetBook(id int) Book {
	sqlStatement := `SELECT title, price FROM books WHERE id=$1;`
	var title string
	var price float32
	row := DB.QueryRow(sqlStatement, id)
	switch err := row.Scan(&title, &price); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(title, price)
	default:
		panic(err)
	}

	b := Book{title, price}
	return b
}
