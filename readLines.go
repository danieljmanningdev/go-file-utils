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

func ReadLines(f string) ([]string, error) {
	file, err := os.Open(string(f))
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
