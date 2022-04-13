package utils

import "os"

func FileExists(path string) bool {
	f, err := os.Stat(path)
	return err == nil && !f.IsDir()
}
