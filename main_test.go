package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBooks(t *testing.T) {
	// Create a new request to the /books endpoint
	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler with the request and recorder
	handler := http.HandlerFunc(getBooks)
	handler.ServeHTTP(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the Content-Type header of the response
	if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want application/json",
			contentType)
	}

	// Check the response body
	var responseBooks []Book
	err = json.NewDecoder(rr.Body).Decode(&responseBooks)
	if err != nil {
		t.Fatal(err)
	}

	if len(responseBooks) != len(books) {
		t.Errorf("handler returned wrong number of books: got %v want %v",
			len(responseBooks), len(books))
	}

	for i, responseBook := range responseBooks {
		if responseBook.Title != books[i].Title {
			t.Errorf("handler returned wrong book title: got %v want %v",
				responseBook.Title, books[i].Title)
		}
		if responseBook.Author != books[i].Author {
			t.Errorf("handler returned wrong book author: got %v want %v",
				responseBook.Author, books[i].Author)
		}
	}
}
