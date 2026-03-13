package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
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

	var searchTerm = r.URL.Query()

	var term string
	var value string
	for k, v := range searchTerm {
		if strings.Contains("author, title, isbn, booksbyauthor", strings.ToLower(k)) {
			term = strings.ToLower(k)
			value = strings.ToLower(strings.Join(v, ""))
		}
	}
	log.Println(term + " " + value)

	if value == "" {
		http.Error(w, "search query parameter is required", http.StatusBadRequest)
		return
	} else {
		fmt.Printf("Search term: %s\n", term)
		switch term {
		case "author", "title", "isbn":
			GetBookByTerm(w, r, term, value)
		case "ListOfBooksByAuthor":
			GetListOfBookByAuthor(w, r)
		default:
			http.Error(w, "Invalid search type", http.StatusBadRequest)
		}
	}
}

func GetBookByTerm(w http.ResponseWriter, r *http.Request, searchTerm string, value string) {
	user := r.Context().Value("user").(User)
	log.Println("Looking for book, ", user.Name)

	var term string
	switch searchTerm {
	case "isbn":
		response, err := searchByISBN(value)
		finish(w, response, err, term)
	case "title":
		response, err := searchByTitle(value)
		finish(w, response, err, term)
	case "author":
		response, err := searchByAuthor(value)
		finish(w, response, err, term)
	default:
		http.Error(w, "Invalid search type", http.StatusBadRequest)
	}

}

func finish[T any](w http.ResponseWriter, response T, err error, term string) {

	if err != nil {
		http.Error(w, "Book(s) using term '"+term+"' not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func GetListOfBookByAuthor(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(User)
	log.Println("Making list, ", user.Name)

	response := getListByAuthor()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

// func GetBookByISBN(w http.ResponseWriter, r *http.Request) {

// 	user := r.Context().Value("user").(User)
// 	log.Println("Looking for book, ", user.Name)

// 	var isbn = r.URL.Query().Get("isbn")
// 	if isbn == "" {
// 		http.Error(w, "ISBN number is required", http.StatusNotFound)
// 		return
// 	}

// 	response, err := searchByISBN(isbn)
// 	if err != nil {
// 		http.Error(w, "Book with ISBN: "+isbn+" not found", http.StatusNotFound)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	json.NewEncoder(w).Encode(response)
// }

// func GetBooksByTitle(w http.ResponseWriter, r *http.Request) {
// 	user := r.Context().Value("user").(User)
// 	log.Println("Looking for book, ", user.Name)

// 	var term = r.URL.Query().Get("title")
// 	if term == "" {
// 		http.Error(w, "Author is required", http.StatusNotFound)
// 		return
// 	}

// 	response, err := searchByTitle(term)
// 	if err != nil {
// 		http.Error(w, "Book by "+term+" not found", http.StatusNotFound)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	json.NewEncoder(w).Encode(response)
// }

// func GetBookByAuthor(w http.ResponseWriter, r *http.Request) {

// 	user := r.Context().Value("user").(User)
// 	log.Println("Looking for book, ", user.Name)

// 	var author = r.URL.Query().Get("author")
// 	if author == "" {
// 		http.Error(w, "Author is required", http.StatusNotFound)
// 		return
// 	}

// 	response, err := searchByAuthor(author)
// 	if err != nil {
// 		http.Error(w, "Book by "+author+" not found", http.StatusNotFound)
// 		return
// 	}

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
