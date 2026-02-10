package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handleClientProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetBook(w, r)
	case http.MethodPatch:
		UpdateBook(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(User)
	log.Println("Looking for book...", user.Name)

	var isbn = r.URL.Query().Get("isbn")
	book, ok := bookDatabase[isbn]
	if !ok || isbn == "" {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	response := Book{
		Title: book.Title,
		Desc:  book.Desc,
		Type:  book.Type,
		ISBN:  book.ISBN,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(User)
	log.Println("Updating book...", user.Name)

	var isbn = r.URL.Query().Get("isbn")
	book, ok := bookDatabase[isbn]
	if !ok || isbn == "" {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Decode the JSON payload directly into struct
	var payloadData Book
	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Update Profile
	book.Title = payloadData.Title
	book.Desc = payloadData.Desc
	book.Type = payloadData.Type
	book.ISBN = payloadData.ISBN
	bookDatabase[book.ISBN] = book

	w.WriteHeader(http.StatusOK)
}
