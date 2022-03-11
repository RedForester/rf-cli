package manifest

import (
	"gopkg.in/yaml.v3"
	"io"
)

func Read(r io.Reader) (*Manifest, error) {
	var manifest Manifest

	d := yaml.NewDecoder(r)
	if err := d.Decode(&manifest); err != nil {
		return nil, err
	}

	return &manifest, nil
}
