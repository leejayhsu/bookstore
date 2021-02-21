package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// DB Create an exported global variable to hold the database connection pool.
var DB *sql.DB

// Book is a model of a book
type Book struct {
	ID    int     `json:"bookId"`
	Title string  `json:"title"`
	Price float32 `json:"price"`
}

// InsertBook inserts a book
func InsertBook(b Book) Book {
	sqlStatement := `
INSERT INTO books (title, price)
VALUES ($1, $2)
RETURNING id`
	id := 0
	err := DB.QueryRow(sqlStatement, b.Title, b.Price).Scan(&id)
	if err != nil {
		fmt.Println("Something went very wrong when creating a book!")
	}
	b.ID = id
	return b
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
		fmt.Println("Something went very wrong when getting a book!")
	}

	b := Book{id, title, price}
	return b
}

// UpdateBook updates a book
func UpdateBook(b Book) {
	sqlStatement := `
UPDATE books
SET title = $2, price = $3
WHERE id = $1;`
	_, err := DB.Exec(sqlStatement, b.ID, b.Title, b.Price)
	if err != nil {
		panic(err)
	}
}

//PatchBook patches a book
func PatchBook(id int, payload map[string]interface{}) Book {
	book := GetBook(id)
	if title, ok := payload["title"]; ok {
		if val, ok := title.(string); ok {
			book.Title = val
		}
	}
	if price, ok := payload["price"]; ok {
		if _, ok := price.(json.Number); ok {
			x, err := price.(json.Number).Float64()
			if err != nil {
				panic("bad conversion to float64")
			}
			book.Price = float32(x)
		}
	}

	UpdateBook(book)
	return book
}
