package gofileutils

import "os"

// OpenFile opens a file from the given path and returns the file pointer.
// Panics if the file cannot be opened.
//
// Author: Daniel Manning <daniel@danieljmanningdev.com>
// Created: 2026
// Last Modified: 28 August 2026
func OpenFile(f Filepath) *os.File {
	file, err := os.Open(string(f))
	if err != nil {
		panic(err)
	}
	return file
}
