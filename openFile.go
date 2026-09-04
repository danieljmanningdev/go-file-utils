package gofileutils

import "os"

// OpenFile opens a file from the given path and returns the file pointer.
// Panics if the file cannot be opened.
//
// Author: Daniel Manning <daniel@danieljmanningdev.com>
// Created: 2026

func OpenFile(f string) (*os.File, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	return file, nil
}
