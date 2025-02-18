package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lokesh1jha/bookstore/pkg/models"
)

// NewBook is a variable to hold the new book
var NewBook models.Book

// GetBooks is a function to get all books
func GetBook(w http.ResponseWriter, r *http.Request) {
	newBook := models.GetAllBooks()
	res, _ := json.Marshal(newBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// GetBook is a function to get a book by ID
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// CreateBook is a function to create a book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	_ = json.NewDecoder(r.Body).Decode(&NewBook)
	NewBook.CreateBook()
	res, _ := json.Marshal(NewBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// UpdateBook is a function to update a book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Fetch the existing book
	existingBook, _ := models.GetBookById(ID)

	// If the book with the given ID doesn't exist, return an error
	if existingBook.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Decode the request body into a new book object
	var newBook models.Book
	err = json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		fmt.Println("Error decoding request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Update only the fields that are provided in the request
	if newBook.Title != "" {
		existingBook.Title = newBook.Title
	}
	if newBook.Author != "" {
		existingBook.Author = newBook.Author
	}
	if newBook.Publication != "" {
		existingBook.Publication = newBook.Publication
	}
	// Add other fields as necessary

	// Save the updated book
	updatedBook := existingBook.UpdateBook(ID)
	res, _ := json.Marshal(updatedBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DeleteBook is a function to delete a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
