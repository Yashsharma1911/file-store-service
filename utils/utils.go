package utils

import (
	"fmt"
	"io"
	"net/http"
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
