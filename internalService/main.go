package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for:", r.URL.Path)
	fmt.Fprintf(w, "Hello, you've hit %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", helloHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
