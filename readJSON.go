package gofileutils

import (
	"encoding/json"
	"os"
)

// ReadJSON opens a file from the given path and decodes its JSON contents
// directly into the provided target variable (typically a pointer to a struct or map).
// Panics if the file cannot be opened or if the JSON decoding fails.
//
// Author: Daniel Manning <daniel@danieljmanningdev.com>
// Created: 2026
// Last Modified: 28 August 2026
func ReadJSON(f filepath, v interface{}) {
	file, err := os.Open(string(f))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(v); err != nil {
		panic(err)
	}
}
