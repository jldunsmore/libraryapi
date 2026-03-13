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
		if strings.Contains("author, title, isbn, booksbyauthor, listofbooksbyauthor", strings.ToLower(k)) {
			term = strings.ToLower(k)
			value = strings.ToLower(strings.Join(v, ""))
		}
	}
	log.Println(term + " " + value)

	if value == "" && term != "listofbooksbyauthor" {
		http.Error(w, "search query parameter is required", http.StatusBadRequest)
		return
	} else {
		fmt.Printf("Search term: %s\n", term)
		switch term {
		case "author", "title", "isbn":
			GetBookByTerm(w, r, term, value)
		case "listofbooksbyauthor":
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
	log.Println("Making list of books by author, " + user.Name)

	response := getListByAuthor()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
