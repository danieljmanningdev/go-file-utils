package gofileutils

import (
	"bufio"
	"os"
)

// WriteLines creates or overwrites a file at the given path, writing each string
// from the slice as a separate line.
//
// Author: Daniel Manning <daniel@danieljmanningdev.com>
// Created: 2026
// Last Modified: 28 August 2026
func WriteLines(f Filepath, lines []string) error {
	file, err := os.Create(string(f))
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
	writer.Flush()
	return nil
}
