package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"New System", 2},
		{"Hello, world!", 2},
		{"", 0},
		{"Go   is awesome", 3},
		{"word\nword", 2},
		{"word\tword", 2},
		{"  multiple  spaces  here  ", 3},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			// Call the CountWords function
			count, err := CountWords(test.input)
			if err != nil {
				t.Fatalf("Expected no error, but got: %v", err)
			}

			assert.Equal(t, test.expected, count)
		})
	}
}
