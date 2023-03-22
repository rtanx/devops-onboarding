package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	// Add some books to our collection
	books = append(books, Book{Title: "The Catcher in the Rye", Author: "J.D. Salinger"})
	books = append(books, Book{Title: "To Kill a Mockingbird", Author: "Harper Lee"})
	books = append(books, Book{Title: "1984", Author: "George Orwell"})

	// Define our endpoint
	http.HandleFunc("/books", getBooks)
	http.HandleFunc("/db", testDBConnection)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode our list of books to JSON
	json.NewEncoder(w).Encode(books)
}

func testDBConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := sql.Open("postgres", "postgres://bookstoreapi:s3cretP@ssword@db/bookstore?sslmode=disable")
	if err != nil {
		json.NewEncoder(w).Encode(map[string]any{
			"sucess":  false,
			"message": fmt.Sprintf("Error opening database: %v", err),
		})
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		json.NewEncoder(w).Encode(map[string]any{
			"sucess":  false,
			"message": fmt.Sprintf("Error connecting to database: %v", err),
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"sucess":  true,
		"message": "connected to databse",
	})
}
