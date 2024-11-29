package main

import (
	"log"

	"github.com/Yashsharma1911/file-store-service/server"
	"github.com/Yashsharma1911/file-store-service/server/dataAccess"
	"github.com/Yashsharma1911/file-store-service/server/database"
	"github.com/Yashsharma1911/file-store-service/server/handlers"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize the MinIO client
	minioClient, err := database.NewMinIOClient()
	if err != nil {
		log.Fatalf("Error initializing MinIO client: %v", err)
	}

	// Create FileDataAccess instance
	fileDataAccess := dataAccess.NewFileDataAccess(minioClient)

	// Create Handlers instance with dependencies injected
	h := handlers.NewHandlers(*fileDataAccess)

	// Initialize Echo web server
	e := echo.New()

	// Middleware for logging and recovering from panics
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Set up routes using SetupRouter function
	server.SetupRouter(e, *h)

	// Start the server on port 8080
	log.Fatal(e.Start(":8080"))
}
