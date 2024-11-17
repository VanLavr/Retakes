package main

import (
	"log"

	"github.com/VanLavr/Retakes/api/gen"
	"github.com/VanLavr/Retakes/internal/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()
	// Create a new server instance
	srv := Server{
		AdminController:   &controller.AdminController{},
		AuthController:    &controller.AuthController{},
		StudentController: &controller.StudentController{},
		TeacherController: &controller.TeacherController{},
	}

	wrapper := &gen.ServerInterfaceWrapper{
		Handler: srv,
	}

	// Use the generated router
	gen.RegisterHandlers(router, wrapper)

	// Start the server
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

type Server struct {
	*controller.AdminController
	*controller.AuthController
	*controller.StudentController
	*controller.TeacherController
}
