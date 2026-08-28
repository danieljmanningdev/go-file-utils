package gofileutils

import "os"

// FileExists checks whether a file exists and is accessible at the given path.
//
// Author: Daniel Manning <daniel@danieljmanningdev.com>
// Created: 2026
// Last Modified: 28 August 2026
func FileExists(f filepath) bool {
	info, err := os.Stat(string(f))
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
