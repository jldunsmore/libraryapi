package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handleClientProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		SearchDatabase(w, r)
	case http.MethodPatch:
		//UpdateBook(w, r)
		fmt.Fprintln(w, "Update book endpoint is under construction.")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func SearchDatabase(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(User)
	log.Println("searching database, ", user.Name)

	var searchTerm = r.URL.Query().Get("search")
	if searchTerm == "" {
		http.Error(w, "search query parameter is required", http.StatusBadRequest)
		return
	} else {
		fmt.Printf("Search term: %s\n", searchTerm)
		switch searchTerm {
		case "BookByISBN":
			GetBookByISBN(w, r)
		case "BooksByTitle":
			//GetBooksByTitle(w, r)
		case "BooksByAuthor":
			//GetBookByAuthor(w, r)
		case "ListOfBooksByAuthor":
			GetListOfBookByAuthor(w, r)
		default:
			http.Error(w, "Invalid search type", http.StatusBadRequest)
		}
	}
}

func GetBookByISBN(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(User)
	log.Println("Looking for book, ", user.Name)

	var isbn = r.URL.Query().Get("isbn")
	response = GetBook_ISBN(isbn)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func GetBookByTitle(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(User)
	log.Println("Looking for book, ", user.Name)

	//var title = r.URL.Query().Get("title")
	//database("Title", title)
}

func GetListOfBookByAuthor(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(User)
	log.Println("Making list, ", user.Name)

	var response = GetListByAuthor()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

// func GetBook(w http.ResponseWriter, r *http.Request) {
// 	user := r.Context().Value("user").(User)
// 	log.Println("Looking for book...", user.Name)

// 	var isbn = r.URL.Query().Get("isbn")
// 	database("ISBN", isbn)

// 	w.Header().Set("Content-Type", "application/json")

// 	json.NewEncoder(w).Encode(response)
// }

// func UpdateBook(w http.ResponseWriter, r *http.Request) {
// 	user := r.Context().Value("user").(User)
// 	log.Println("Updating book...", user.Name)

// 	var isbn = r.URL.Query().Get("isbn")
// 	book, ok := bookDatabase[isbn]
// 	if !ok || isbn == "" {
// 		http.Error(w, "Book not found", http.StatusNotFound)
// 		return
// 	}

// 	// Decode the JSON payload directly into struct
// 	var payloadData Book
// 	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
// 		http.Error(w, "Invalid JSON", http.StatusBadRequest)
// 		return
// 	}
// 	defer r.Body.Close()

// 	// Update Book
// 	if payloadData.Title != "" {
// 		book.Title = payloadData.Title
// 	}
// 	if payloadData.Desc != "" {
// 		book.Desc = payloadData.Desc
// 	}
// 	if payloadData.Type != "" {
// 		book.Type = payloadData.Type
// 	}
// 	if payloadData.ISBN != "" {
// 		book.ISBN = isbn
// 	}
// 	bookDatabase[book.ISBN] = book

// 	fmt.Printf("Book updated: %+v\n", book)

// 	w.WriteHeader(http.StatusOK)
// }
