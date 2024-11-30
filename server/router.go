package server

import (
	"github.com/Yashsharma1911/file-store-service/server/handlers"
	"github.com/labstack/echo/v4"
)

/**
* Setup routers and middlewares here
 */
func SetupRouter(e *echo.Echo, h handlers.Handlers) {
	// Group API routes under "/api"
	api := e.Group("/api")

	// File routes
	api.POST("/files", h.AddFile)
	api.GET("/files", h.ListFiles)
	api.GET("/files/:name", h.GetFile)
	api.DELETE("/files/:name", h.RemoveFile)
	api.PUT("/files/:name", h.UpdateFile)

	// Word count and frequency routes
	api.GET("/wc", h.WordCount)
	api.GET("/frequent", h.MostFrequentWords)
}
