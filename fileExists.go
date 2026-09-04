package gofileutils

import "os"

// FileExists checks whether a file exists and is accessible at the given path.
//
// Author: Daniel Manning <daniel@danieljmanningdev.com>
// Created: 2026
// Last Modified: 28 August 2026
func FileExists(f string) bool {
	info, err := os.Stat(string(f))
	if err != nil {
		return false
	}

	return !info.IsDir()
}
