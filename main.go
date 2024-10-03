package main

import (
	"log"
	"net/http"

	"triples/http_utils"
)

func main() {
	http.HandleFunc("/", http_utils.Handler)
	address := ":8091"
	log.Printf("Starting development server at http://localhost%s/\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
