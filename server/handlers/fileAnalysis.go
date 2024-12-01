// This file contains analysis handlers for operations on stored files

package handlers

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
	"sync"
	"unicode"

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

// MostFrequentWords processes files to find the most/least frequent words.
func (h *Handlers) MostFrequentWords(c echo.Context) error {
	ctx := c.Request().Context()

	files, _ := h.FileDataAccess.ListFiles(ctx)
	limit, order := utils.DefaultParams(c)
	// Word frequency map
	wordFreq := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go func(file *models.FileMetadata) {
			defer wg.Done()

			fileContent, err := h.FileDataAccess.GetFileContent(ctx, file.FileName)
			if err != nil {
				fmt.Printf("Error fetching file content for %s: %v\n", file.FileName, err)
				return
			}

			// Tokenize words and update word frequencies
			// Tokenize text based on letters and numbers while ignoring any non-letter, non-number characters
			words := strings.FieldsFunc(string(fileContent), func(r rune) bool {
				return !unicode.IsLetter(r) && !unicode.IsNumber(r)
			})

			// Local frequency is map to hold current goroutine processed data
			// this is to ensure each goroutine is independent and does not wait for wordFreq (Global map) to get free
			localFreq := make(map[string]int)
			for _, word := range words {
				normalizedWord := strings.ToLower(word)
				localFreq[normalizedWord]++
			}

			mu.Lock()
			for word, count := range localFreq {
				wordFreq[word] += count
			}
			mu.Unlock()
		}(file)
	}

	wg.Wait()

	// Convert word frequency map to a slice
	freqSlice := make([]models.Word, 0, len(wordFreq))

	for word, count := range wordFreq {
		freqSlice = append(freqSlice, struct {
			Word  string
			Count int
		}{Word: word, Count: count})
	}

	// Sort based on frequency
	sort.Slice(freqSlice, func(i, j int) bool {
		if order == "asc" {
			return freqSlice[i].Count < freqSlice[j].Count
		}
		return freqSlice[i].Count > freqSlice[j].Count
	})

	// Limit results
	if limit < len(freqSlice) {
		freqSlice = freqSlice[:limit]
	}

	return c.JSON(http.StatusOK, freqSlice)
}
