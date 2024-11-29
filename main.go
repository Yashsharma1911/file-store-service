package main

import (
	"log"

	"github.com/Yashsharma1911/file-store-service/server"
	"github.com/Yashsharma1911/file-store-service/server/dataAccess"
	"github.com/Yashsharma1911/file-store-service/server/database"
	"github.com/Yashsharma1911/file-store-service/server/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize the MinIO client
	minioClient, err := database.NewMinIOClient()
	if err != nil {
		log.Fatalf("Error initializing MinIO client: %v", err)
	}
	// Initialize DataAccess Instance
	fileDataAccess := dataAccess.NewFileDataAccess(minioClient)

	// Initialize Handlers instance with dependencies injected
	h := handlers.NewHandlers(*fileDataAccess)

	// start echo server
	e := echo.New()

	// Set up routes using SetupRouter function
	server.SetupRouter(e, *h)

	// Start the server on port 8080
	log.Fatal(e.Start(":8080"))
}
