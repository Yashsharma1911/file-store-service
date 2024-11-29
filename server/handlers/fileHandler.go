package handlers

import (
	"net/http"

	"github.com/Yashsharma1911/file-store-service/server/models"
	"github.com/labstack/echo/v4"
)

// AddFile handles the uploading of files
func (h *Handlers) AddFile(c echo.Context) error {
	// Create a new file metadata struct
	file := new(models.FileMetadata)

	// Get the file from the request
	// Assuming the file is sent as a multipart form-data
	fileData, err := c.FormFile("file") // "file" is the form field name
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid file"})
	}

	// Open the file from the multipart form data
	fileContent, err := fileData.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not open file"})
	}
	defer fileContent.Close()

	// Read the file into memory
	fileBytes := make([]byte, fileData.Size)
	_, err = fileContent.Read(fileBytes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not read file"})
	}

	// Upload the file using FileDataAccess
	uploadedFile, err := h.FileDataAccess.UploadFile(c.Request().Context(), fileBytes, fileData.Filename)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Set metadata for the uploaded file
	file.FileName = uploadedFile.FileName
	file.Size = uploadedFile.Size

	// Return the file metadata as a JSON response
	return c.JSON(http.StatusCreated, file)
}

// GetFile handles fetching file metadata by name
func (h *Handlers) GetFile(c echo.Context) error {
	// Get the file name from the URL parameter
	name := c.Param("name")

	// Retrieve the file metadata from FileDataAccess
	fileMetadata, err := h.FileDataAccess.GetFileMetadata(c.Request().Context(), name)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	// Return the file metadata as a JSON response
	return c.JSON(http.StatusOK, fileMetadata)
}

// ListFiles handles fetching a list of files
func (h *Handlers) ListFiles(c echo.Context) error {
	// Get the list of files using FileDataAccess
	files, err := h.FileDataAccess.ListFiles(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Return the list of files as a JSON response
	return c.JSON(http.StatusOK, files)
}

// RemoveFile handles file deletion by name
func (h *Handlers) RemoveFile(c echo.Context) error {
	// Get the file name from the URL parameter
	name := c.Param("name")

	// Attempt to delete the file using FileDataAccess
	err := h.FileDataAccess.DeleteFile(c.Request().Context(), name)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	// Return a success message
	return c.JSON(http.StatusOK, map[string]string{"message": "file deleted"})
}

// UpdateFile handles updating an existing file by replacing it with a new one
func (h *Handlers) UpdateFile(c echo.Context) error {
	// Get the file name from the URL parameter
	oldFileName := c.Param("name")

	// Get the new file from the request
	// Assuming the new file is sent as a multipart form-data
	fileData, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid file"})
	}

	// Open the new file from the multipart form data
	fileContent, err := fileData.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not open new file"})
	}
	defer fileContent.Close()

	// Read the file into memory
	fileBytes := make([]byte, fileData.Size)
	_, err = fileContent.Read(fileBytes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not read new file"})
	}

	// Update the file using FileDataAccess (deletes old and uploads new)
	uploadedFile, err := h.FileDataAccess.UpdateFile(c.Request().Context(), oldFileName, fileData.Filename, fileBytes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Return the updated file metadata as a JSON response
	return c.JSON(http.StatusOK, uploadedFile)
}
