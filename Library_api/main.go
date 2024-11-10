package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/books", getBooks)
	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Simple REST API in Go!")
}

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func init() {
	books = []Book{
		{ID: "1", Title: "1984", Author: "George Orwell"},
		{ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee"},
		{ID: "3", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
	}
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
