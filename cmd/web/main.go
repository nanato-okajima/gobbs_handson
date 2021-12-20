package main

import (
	"log"
	"net/http"
)

func main() {
	mux := routes()

	log.Println("Starting web server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
