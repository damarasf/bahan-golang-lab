package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	PORT := ":8081"
	http.Handle("/", http.FileServer(http.Dir("polymer")))
	fmt.Println("Listening on", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
