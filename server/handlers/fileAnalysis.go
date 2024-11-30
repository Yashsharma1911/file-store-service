// This file contains analysis handlers for operations on stored files

package handlers

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/Yashsharma1911/file-store-service/server/models"
	"github.com/Yashsharma1911/file-store-service/utils"
	"github.com/labstack/echo/v4"
)

// WordCount calculates the total word count across all files stored in the MinIO server.
func (h *Handlers) WordCount(c echo.Context) error {
	var totalWords int
	var mu sync.Mutex
	var wg sync.WaitGroup
	var errors []string

	// Fetch the list of files
	files, err := h.FileDataAccess.ListFiles(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("error fetching file list: %v", err),
		})
	}

	// Launch goroutines to do parallel operations to count total words in stored files
	for _, file := range files {
		wg.Add(1)
		go func(file *models.FileMetadata) {
			defer wg.Done()

			// Fetch file content
			fileContent, err := h.FileDataAccess.GetFileContent(c.Request().Context(), file.FileName)
			if err != nil {
				mu.Lock()
				errors = append(errors, fmt.Sprintf("error fetching file %s: %v", file.FileName, err))
				mu.Unlock()
				return
			}

			// Count words
			words, _ := utils.CountWords(fileContent)
			mu.Lock()
			totalWords += words
			mu.Unlock()
		}(file)
	}

	wg.Wait()

	if len(errors) > 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "some files could not be processed",
			"details": errors,
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("Total word count is %d", totalWords),
	})
}
