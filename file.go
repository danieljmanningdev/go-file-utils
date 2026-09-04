//
// Daniel J. Manning
//
// Copyright © 2026 Daniel J. Manning.
// SPDX-License-Identifier: MIT
//

package gofileutils

import (
	"bufio"
	"os"
)

// FileExists checks whether a file exists and is accessible at the given path.
func FileExists(f string) bool {
	info, err := os.Stat(f)
	if err != nil {
		return false
	}

	return !info.IsDir()
}

// OpenFile opens a file from the given path and returns the file pointer.
// Panics if the file cannot be opened.
func OpenFile(f string) (*os.File, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// ReadLines opens a file from the given path, reads it line by line using a scanner,
// appends each line to a slice of strings, and returns the full collection.
// It handles file closing automatically via defer and panics if any I/O errors occur.
func ReadLines(f string) ([]string, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// WriteLines creates or overwrites a file at the given path, writing each string
// from the slice as a separate line.
func WriteLines(f string, lines []string) error {
	file, err := os.Create(f)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	if err := writer.Flush(); err != nil {
		return err
	}

	return nil
}
