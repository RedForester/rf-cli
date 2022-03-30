package manifest

import (
	"github.com/deissh/rf-cli/internal/utils"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

func Read(r io.Reader) (*Manifest, error) {
	var manifest Manifest

	d := yaml.NewDecoder(r)
	if err := d.Decode(&manifest); err != nil {
		return nil, err
	}

	return &manifest, nil
}

// ReadByPath file and return Manifest or error
func ReadByPath(path string) (*Manifest, error) {
	if !utils.FileExists(path) {
		return nil, os.ErrNotExist
	}

	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return Read(f)
}
