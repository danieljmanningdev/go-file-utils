//
// Daniel J. Manning
//
// Copyright © 2026 Daniel J. Manning.
// SPDX-License-Identifier: MIT
//

package gofileutils

import (
	"encoding/json"
	"os"
)

// ReadJSON opens a file from the given path and decodes its JSON contents
// directly into the provided target variable (typically a pointer to a struct or map).
// Panics if the file cannot be opened or if the JSON decoding fails.
func ReadJSON(f string, v any) error {
	file, err := os.Open(f)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(v); err != nil {
		return err
	}

	return nil
}

// WriteJSON encodes any given data structure into pretty-printed JSON
// and writes it to a file at the specified path.
func WriteJSON(f string, v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(f, data, 0644); err != nil {
		return err
	}
	return nil
}
