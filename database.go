package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
)

// type Title struct {
// 	Title string `json:"title"`
// 	Id    string `json:"id"`
// }

type BookAuthors struct {
	Authors []string `json:"authors"`
	ISBN    string   `json:"isbn"`
}

type Book struct {
	Kind       string `json:"kind"`
	ID         string `json:"id"`
	Etag       string `json:"etag"`
	SelfLink   string `json:"selfLink"`
	VolumeInfo struct {
		Title               string   `json:"title"`
		Authors             []string `json:"authors"`
		PublishedDate       string   `json:"publishedDate"`
		Description         string   `json:"description"`
		IndustryIdentifiers []struct {
			Type       string `json:"type"`
			Identifier string `json:"identifier"`
		} `json:"industryIdentifiers"`
		ReadingModes struct {
			Text  bool `json:"text"`
			Image bool `json:"image"`
		} `json:"readingModes"`
		PageCount           int    `json:"pageCount"`
		PrintType           string `json:"printType"`
		MaturityRating      string `json:"maturityRating"`
		AllowAnonLogging    bool   `json:"allowAnonLogging"`
		ContentVersion      string `json:"contentVersion"`
		PanelizationSummary struct {
			ContainsEpubBubbles  bool `json:"containsEpubBubbles"`
			ContainsImageBubbles bool `json:"containsImageBubbles"`
		} `json:"panelizationSummary"`
		ImageLinks struct {
			SmallThumbnail string `json:"smallThumbnail"`
			Thumbnail      string `json:"thumbnail"`
		} `json:"imageLinks"`
		Language            string `json:"language"`
		PreviewLink         string `json:"previewLink"`
		InfoLink            string `json:"infoLink"`
		CanonicalVolumeLink string `json:"canonicalVolumeLink"`
	} `json:"volumeInfo"`
	SaleInfo struct {
		Country     string `json:"country"`
		Saleability string `json:"saleability"`
		IsEbook     bool   `json:"isEbook"`
	} `json:"saleInfo"`
	AccessInfo struct {
		Country                string `json:"country"`
		Viewability            string `json:"viewability"`
		Embeddable             bool   `json:"embeddable"`
		PublicDomain           bool   `json:"publicDomain"`
		TextToSpeechPermission string `json:"textToSpeechPermission"`
		Epub                   struct {
			IsAvailable bool `json:"isAvailable"`
		} `json:"epub"`
		Pdf struct {
			IsAvailable bool `json:"isAvailable"`
		} `json:"pdf"`
		WebReaderLink       string `json:"webReaderLink"`
		AccessViewStatus    string `json:"accessViewStatus"`
		QuoteSharingAllowed bool   `json:"quoteSharingAllowed"`
	} `json:"accessInfo"`
	SearchInfo struct {
		TextSnippet string `json:"textSnippet"`
	} `json:"searchInfo"`
}

type BookCard struct {
	Title string "json:\"type\""
	ISBN  []struct {
		Type       string "json:\"type\""
		Identifier string "json:\"identifier\""
	} `json:"industryIdentifiers"`
}

type BookList struct {
	Author string
	List   []BookCard
}

func database() []Book {
	bookJsonFile, err := os.Open("books.json")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully opened books.json")

	defer bookJsonFile.Close()

	byteValue, _ := io.ReadAll(bookJsonFile)

	var books struct {
		Books []Book `json:"items"`
	}

	json.Unmarshal(byteValue, &books)

	return books.Books
}

func searchByISBN(isbn string) (Book, error) {
	var books = database()
	fmt.Printf("Searching for ISBN: %s\n", isbn)

	var response Book
	for i := 0; i < len(books); i++ {
		ISBN := books[i].VolumeInfo.IndustryIdentifiers

		for j := 0; j < len(ISBN); j++ {
			if string(ISBN[j].Identifier) == isbn {
				return books[i], nil
			}
		}
	}
	return response, errors.New("404")
}

func searchByTitle(title string) ([]BookCard, error) {
	var books = database()
	fmt.Printf("Searching for Author: %s\n", title)

	var response []BookCard
	for i := 0; i < len(books); i++ {
		bookTitle := books[i].VolumeInfo.Title

		if strings.Contains(strings.ToLower(string(bookTitle)), strings.ToLower(title)) {
			book := BookCard{
				Title: books[i].VolumeInfo.Title,
				ISBN:  books[i].VolumeInfo.IndustryIdentifiers,
			}
			response = append(response, book)
		}
	}
	if len(response) > 0 {
		return response, nil
	} else {
		return response, errors.New("404")
	}
}

func searchByAuthor(author string) (BookList, error) {
	var books = database()
	fmt.Printf("Searching for Author: %s\n", author)

	var response BookList
	for i := 0; i < len(books); i++ {
		authors := books[i].VolumeInfo.Authors

		for j := 0; j < len(authors); j++ {
			if strings.EqualFold(string(authors[j]), author) {
				book := BookCard{
					Title: books[i].VolumeInfo.Title,
					ISBN:  books[i].VolumeInfo.IndustryIdentifiers,
				}
				response.List = append(response.List, book)
			}
		}
	}
	if len(response.List) > 0 {
		response.Author = author
		return response, nil
	} else {
		return response, errors.New("404")
	}
}

func getListByAuthor() []BookList {
	var books = database()

	var authors []BookList
	var authorMap = make(map[string][]BookCard)

	log.Println(len(books))
	for i := 0; i < len(books); i++ {
		bookAuthors := books[i].VolumeInfo.Authors
		log.Println("searching for: ", bookAuthors)
		// list of authors
		for j := 0; j < len(bookAuthors); j++ {
			authorname := bookAuthors[j]
			//does author exsit in map
			//if _, ok := authorMap[authorname]; !ok {
			log.Println("searching for: ", authorname)
			//find books with that author
			for k := 0; k < len(books); k++ {
				if slices.Contains(books[k].VolumeInfo.Authors, authorname) {
					book := BookCard{
						Title: books[k].VolumeInfo.Title,
						ISBN:  books[k].VolumeInfo.IndustryIdentifiers,
					}
					// have we added books with this author
					if _, ok := authorMap[authorname]; ok {
						authorMap[authorname] = append(authorMap[authorname], book)
					} else {
						var list []BookCard
						list = append(list, book)
						authorMap[authorname] = list
					}
				}
			}
			//}
		}
	}
	for k, v := range authorMap {
		authors = append(authors, BookList{Author: k, List: v})
	}
	//fmt.Printf("Authors: %+v\n", authors)
	return authors
}
