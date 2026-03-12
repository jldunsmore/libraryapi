package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Title struct {
	Title string `json:"title"`
	Id    string `json:"id"`
}

type BookAuthors struct {
	Authors []string `json:"authors"`
	ISBN    string   `json:"isbn"`
}

type Author struct {
	Name string
	ISBN []string
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

type Books struct {
	Books []Book `json:"items"`
}

type BookCard struct {
	Title string "json:\"type\""
	ISBN  []struct {
		Type       string "json:\"type\""
		Identifier string "json:\"identifier\""
	} `json:"industryIdentifiers"`
}

type AuthorBookList struct {
	Name     string
	BookList []BookCard
}

// func GetBooksByTitle(Title string) {
// }

// func GetBooksByAuthor(author string) {
// }

func database() Books {
	bookJsonFile, err := os.Open("books.json")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully opened books.json")

	defer bookJsonFile.Close()

	byteValue, _ := io.ReadAll(bookJsonFile)

	var books Books

	json.Unmarshal(byteValue, &books)

	return books
}

func searchByISBN(isbn string) (Book, error) {
	var books = database()
	fmt.Printf("Searching for ISBN: %s\n", isbn)

	var response Book
	for i := 0; i < len(books.Books); i++ {
		bookItem := BookAuthors{
			//Authors: books.Books[i].VolumeInfo.Authors,
			ISBN: books.Books[i].VolumeInfo.IndustryIdentifiers[0].Identifier,
		}

		for j := 0; j < len(bookItem.ISBN); j++ {
			if string(bookItem.ISBN[j]) == isbn {
				return books.Books[i], nil
			}
		}
	}
	return response, errors.New("404")
}

// func searchByTitle(books Books, title string) {
// 	fmt.Printf("Searching for title: %s\n", title)
// }

func searchByAuthor(author string) (AuthorBookList, error) {
	var books = database()
	fmt.Printf("Searching for Author: %s\n", author)

	var response AuthorBookList
	for i := 0; i < len(books.Books); i++ {
		authors := books.Books[i].VolumeInfo.Authors

		for j := 0; j < len(authors); j++ {
			if strings.EqualFold(string(authors[j]), author) {
				book := BookCard{
					Title: books.Books[i].VolumeInfo.Title,
					ISBN:  books.Books[i].VolumeInfo.IndustryIdentifiers,
				}
				response.BookList = append(response.BookList, book)
			}
		}
	}
	if len(response.BookList) > 0 {
		response.Name = author
		return response, nil
	} else {
		return response, errors.New("404")
	}
}

func getListByAuthor() []Author {
	var books = database()

	var authors []Author
	for i := 0; i < len(books.Books); i++ {
		authorList := BookAuthors{
			Authors: books.Books[i].VolumeInfo.Authors,
			ISBN:    books.Books[i].VolumeInfo.IndustryIdentifiers[0].Identifier,
		}
		for j := 0; j < len(authorList.Authors); j++ {
			authorname := authorList.Authors[j]
			authorExists := false
			for k := 0; k < len(authors); k++ {
				if authors[k].Name == authorname {
					authors[k].ISBN = append(authors[k].ISBN, authorList.ISBN)
					authorExists = true
					break
				}
			}
			if !authorExists {
				newAuthor := Author{
					Name: authorname,
					ISBN: []string{authorList.ISBN},
				}
				authors = append(authors, newAuthor)
			}
		}
	}
	//fmt.Printf("Authors: %+v\n", authors)
	return authors
}
