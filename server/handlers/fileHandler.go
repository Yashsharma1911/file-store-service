// This file contains CRUD operations on file
//
// Add | Update | Delete | List All Files
package handlers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"sync"

	"github.com/Yashsharma1911/file-store-service/server/models"
	"github.com/labstack/echo/v4"
)

func (h *Handlers) AddFile(c echo.Context) error {
	filesData, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid files"})
	}

	var uploadedFiles []models.FileMetadata
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, fileHeader := range filesData.File["files"] {
		wg.Add(1)
		go func(c echo.Context, fileHeader *multipart.FileHeader) {
			defer wg.Done()

			existingFile, err := h.FileDataAccess.GetFileMetadata(c.Request().Context(), fileHeader.Filename)
			if err == nil && existingFile != nil {
				c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("%s file already exist", fileHeader.Filename)})
				return
			}

			fileContent, err := fileHeader.Open()
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not open file"})
				return
			}
			defer fileContent.Close()

			fileBytes := make([]byte, fileHeader.Size)
			_, err = fileContent.Read(fileBytes)
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not read file"})
				return
			}

			uploadedFile, err := h.FileDataAccess.UploadFile(c.Request().Context(), fileBytes, fileHeader.Filename)
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
				return
			}

			mu.Lock()
			uploadedFiles = append(uploadedFiles, models.FileMetadata{
				FileName: uploadedFile.FileName,
				Size:     uploadedFile.Size,
			})
			mu.Unlock()

		}(c, fileHeader)
	}

	wg.Wait()
	return c.JSON(http.StatusCreated, uploadedFiles)
}

// GetFile handles fetching file metadata by name
func (h *Handlers) GetFile(c echo.Context) error {
	name := c.Param("name")
	fileMetadata, err := h.FileDataAccess.GetFileMetadata(c.Request().Context(), name)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, fileMetadata)
}

// ListFiles handles fetching all files
func (h *Handlers) ListFiles(c echo.Context) error {
	files, err := h.FileDataAccess.ListFiles(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, files)
}

// RemoveFile handles file deletion by name
func (h *Handlers) RemoveFile(c echo.Context) error {
	name := c.Param("name")
	err := h.FileDataAccess.DeleteFile(c.Request().Context(), name)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "file deleted"})
}

// UpdateFile handles updating an existing file by replacing it with a new one
func (h *Handlers) UpdateFile(c echo.Context) error {
	oldFileName := c.Param("name")

	fileData, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid file"})
	}

	fileContent, err := fileData.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not open new file"})
	}
	defer fileContent.Close()

	fileBytes := make([]byte, fileData.Size)
	_, err = fileContent.Read(fileBytes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not read new file"})
	}

	uploadedFile, err := h.FileDataAccess.UpdateFile(c.Request().Context(), oldFileName, fileData.Filename, fileBytes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, uploadedFile)
}
