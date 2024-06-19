package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrPtr(t *testing.T) {

	str := "String to test for pointer"
	ptrStr := StrPtr(str)
	assert.Equal(t, ptrStr, &str)
	assert.EqualValues(t, ptrStr, &str)
}

func TestGetMimeType(t *testing.T) {
	tests := []struct {
		filename string
		expected string
	}{
		{"document.pdf", "application/pdf"},
		{"image.png", "image/png"},
		{"archive.zip", "application/zip"},
		{"unknownfile.unknown", ""},
		{"text.txt", "text/plain; charset=utf-8"},
		{"no_extension", ""},
	}

	for _, test := range tests {
		t.Run(test.filename, func(t *testing.T) {
			result := GetMimeType(test.filename)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestGetMimeTypeEdgeCases(t *testing.T) {
	t.Run("unknown extension", func(t *testing.T) {
		filename := "file.unknownext"
		expected := ""
		result := GetMimeType(filename)
		assert.Equal(t, expected, result)
	})

	t.Run("empty filename", func(t *testing.T) {
		filename := ""
		expected := ""
		result := GetMimeType(filename)
		assert.Equal(t, expected, result)
	})
}
