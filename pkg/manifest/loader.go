package manifest

import (
	"github.com/deissh/rf-cli/pkg/extension"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

// TODO: extension.Extension -> Manifest struct

func ReadFromFile(file string) (*extension.Extension, error) {
	var data extension.Extension

	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// prepare env vars
	content := []byte(os.ExpandEnv(string(yamlFile)))

	err = yaml.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	return &data, err
}
