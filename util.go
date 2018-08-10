package main

import "os"

// IsDir returns whether the given path points to a directory.
func IsDir(path string) bool {
	file, err := os.Stat(path)
	switch {
	case err != nil:
		return false
	case file.IsDir():
		return true
	default:
		return false
	}
}
