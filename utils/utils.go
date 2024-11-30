package utils

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// MakeRequest is a generic function to make HTTP requests
func MakeRequest(method, url string, body io.Reader, contentType string) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return respBody, nil
}

// CountWords counts the words in string
// Spaces does not count
//
// Example "New System"
// return 2
func CountWords(fileContent string) (int, error) {
	if fileContent == "" {
		return 0, nil
	}

	// Split content by spaces and count words
	words := strings.Fields(fileContent)

	return len(words), nil
}

// DefaultParams return default parameters
// [limit | 10]
// [order | "asc"/"dsc"]
func DefaultParams(c echo.Context) (int, string) {
	limit := 10
	order := "dsc"
	if l := c.QueryParam("limit"); l != "" {
		parsedLimit, err := strconv.Atoi(l)
		if err != nil || parsedLimit <= 0 {
			return 10, "dsc"
		}
		limit = parsedLimit
	}

	if o := c.QueryParam("order"); o != "" {
		if o != "asc" && o != "dsc" {
			return 10, "dsc"
		}
		order = o
	}
	return limit, order
}
