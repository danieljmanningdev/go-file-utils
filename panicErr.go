package gofileutils

// Author: Daniel Manning <daniel@danieljmanningdev.com>
// Created: 2026

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
