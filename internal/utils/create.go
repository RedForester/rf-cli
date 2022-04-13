package utils

import "os"

func CreateFileAndBackup(path string) error {
	const perm = 0o700

	if FileExists(path) {
		if err := os.Rename(path, path+".bkp"); err != nil {
			return err
		}
	}
	_, err := os.OpenFile(path, os.O_CREATE, perm)

	return err
}
