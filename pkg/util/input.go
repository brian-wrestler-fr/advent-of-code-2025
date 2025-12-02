package util

import (
	"os"
	"strings"
)

// ReadInput returns the entire file as a string, trimmed of trailing whitespace
func ReadInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(data))
}

// ReadLines reads a file and splits it by newlines
func ReadLines(path string) []string {
	input := ReadInput(path)
	return strings.Split(input, "\n")
}
