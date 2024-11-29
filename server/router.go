package server

import (
	"github.com/Yashsharma1911/file-store-service/server/handlers"
	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo, h handlers.Handlers) {
	e.POST("/files", h.AddFile)
	e.GET("/files", h.ListFiles)
	e.DELETE("/files/:name", h.RemoveFile)
}
