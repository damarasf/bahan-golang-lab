package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 9000

	http.Handle("/", http.FileServer(http.Dir("polymer-starter-kit")))
	fmt.Printf("Starting server on port %d", port)
	fmt.Printf("http://localhost:9000")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}