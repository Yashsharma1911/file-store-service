package handlers

import (
	"github.com/Yashsharma1911/file-store-service/server/dataAccess"
)

// Handlers struct which will include all handler functions, including file handling
type Handlers struct {
	FileDataAccess dataAccess.FileDataAccess
}

// NewHandlers creates a new instance of Handlers with dependencies injected (e.g., FileDataAccess)
func NewHandlers(fileDataAccess dataAccess.FileDataAccess) *Handlers {
	return &Handlers{FileDataAccess: fileDataAccess}
}
