package main

import (
	"github.com/alvarezjulia/jwt-in-go/cmd/handlers"
	"log"
	"net/http"
)

func main() {
	// "Signin" and "Welcome" are the handlers that we will implement
	http.HandleFunc("/signin", handlers.Signin)
	http.HandleFunc("/welcome", handlers.Welcome)
	http.HandleFunc("/refresh", handlers.Refresh)

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}
