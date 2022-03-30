package manifest

import (
	"fmt"
	"github.com/deissh/rf-cli/internal/utils"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

// Write manifest to given writer or return error
func Write(w io.Writer, info *Manifest) error {
	data, err := yaml.Marshal(info)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = w.Write(data)

	return err
}

// WriteByPath file
// before write truncate file
func WriteByPath(path string, info *Manifest) error {
	if !utils.FileExists(path) {
		return os.ErrNotExist
	}

	if err := os.Truncate(path, 0); err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return Write(f, info)
}
