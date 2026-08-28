package gofileutils

import (
	"bufio"
	"os"
)

// ReadLines opens a file from the given path, reads it line by line using a scanner,
// appends each line to a slice of strings, and returns the full collection.
// It handles file closing automatically via defer and panics if any I/O errors occur.
//
// Author: Daniel Manning <daniel@danieljmanningdev.com>
// Created: 2026
// Last Modified: 2026-08-28
func ReadLines(f Filepath) []string {
	// Open the file using the underlying string value of our custom type
	file, err := os.Open(string(f))
	if err != nil {
		panic(err)
	}
	// Ensure the file is safely closed when the function finishes executing
	defer file.Close()

	var lines []string

	// Create a new scanner to stream the file contents line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for any errors that occurred during the scanning process
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}
