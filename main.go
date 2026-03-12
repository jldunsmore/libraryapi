package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

var middleware = []Middleware{
	TokenAuthMiddleware,
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var handler http.HandlerFunc = handleClientProfile

	for _, middleware := range middleware {
		handler = middleware(handler)
	}

	http.HandleFunc("/library/book", handler)

	server := &http.Server{Addr: ":8080", Handler: handler}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("\nServer error: %v\n", err)
		}
	}()
	fmt.Println("\nServer started on :8080")

	<-ctx.Done()
	fmt.Println("\nrecevied interupt signal")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {

	} else {
		fmt.Println("\ngracefull shutdown complete")
	}
}

// curl --request GET \
// --url http:/localhost:6060/api/messages/public \
// --header 'authorization: Bearer AUTH0-ACCESS-TOKEN'
