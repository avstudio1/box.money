// Node: 1
// Path: /api/main.go
// Purpose: Entrypoint for the backend service, responsible for initializing servers and dependencies.

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Goldstream API is running.")
	})

	log.Println("Starting Goldstream API server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
