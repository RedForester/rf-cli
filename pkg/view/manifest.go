package view

import (
	"fmt"
	"github.com/deissh/rf-cli/pkg/manifest"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type Manifest struct {
	data   *manifest.Manifest
	writer io.Writer
}

func NewManifest(data *manifest.Manifest) *Manifest {
	return &Manifest{
		data:   data,
		writer: os.Stdout,
	}
}

func (m Manifest) Render() error {
	data, err := yaml.Marshal(m.data)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(m.writer, string(data))
	return err
}
