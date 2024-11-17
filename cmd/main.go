package main

import (
	"log"

	"github.com/VanLavr/Retakes/api/gen"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Use the generated router
	gen.RegisterHandlers(router)

	// Start the server
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
