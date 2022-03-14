package manifest

import (
	"fmt"
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

	fmt.Println(data)
	_, err = w.Write(data)

	return err
}
