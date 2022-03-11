package view

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
)

func RenderJSON(writer io.Writer, data interface{}) error {
	res, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(writer, string(res))
	return err
}

func RenderPrettyJSON(writer io.Writer, data interface{}) error {
	res, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(writer, string(res))
	return err
}

func RenderYAML(writer io.Writer, data interface{}) error {
	res, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(writer, string(res))
	return err
}
