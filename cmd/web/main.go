package main

import (
	"log"
	"net/http"

	"gobbs_handson/bulletin-board/internal/handlers"
)

func main() {
	mux := routes()

	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()

	log.Println("Starting web server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
