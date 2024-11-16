package main

import (
	"log"
	"net/http"

	"github.com/VanLavr/Retakes/api/swagger/gen"
)

func main() {
	// Create a new server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: gen.Handler(), // Use the generated router
	}

	log.Println("Server listening on port 8080")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}
