package main

import (
	"log"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

var middleware = []Middleware{
	TokenAuthMiddleware,
}

func main() {
	var handler http.HandlerFunc = handleClientProfile

	for _, middleware := range middleware {
		handler = middleware(handler)
	}

	http.HandleFunc("/library/book", handler)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

// curl --request GET \
// --url http:/localhost:6060/api/messages/public \
// --header 'authorization: Bearer AUTH0-ACCESS-TOKEN'
