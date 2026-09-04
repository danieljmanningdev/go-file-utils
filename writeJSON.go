package gofileutils

import (
	"encoding/json"
	"os"
)

// WriteJSON encodes any given data structure into pretty-printed JSON
// and writes it to a file at the specified path.
//
// Author: Daniel Manning <daniel@danieljmanningdev.com>

func WriteJSON(f string, v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(string(f), data, 0644); err != nil {
		return err
	}
	return nil
}
