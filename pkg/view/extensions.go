package view

import (
	"encoding/json"
	"fmt"
	"github.com/deissh/rf-cli/pkg/rf"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"strings"
	"text/tabwriter"
)

type ExtensionList struct {
	data   *[]rf.Extension
	writer io.Writer
}

func NewExtensionList(data *[]rf.Extension) *ExtensionList {
	return &ExtensionList{
		data:   data,
		writer: os.Stdout,
	}
}

func (l ExtensionList) header() []string {
	return []string{
		"ID",
		"NAME",
		"EMAIL",
		"BASE URL",
		"DESCRIPTION",
	}
}

func (l ExtensionList) printHeader(wr io.Writer) {
	head := strings.Join(l.header(), "\t")
	_, _ = fmt.Fprintln(wr, head)
}

func (l ExtensionList) Render() error {
	wr := tabwriter.NewWriter(l.writer, 0, 0, 2, ' ', 0)

	l.printHeader(wr)

	for _, e := range *l.data {
		baseUrl := ""
		if e.BaseURL != nil {
			baseUrl = *e.BaseURL
		}

		line := []string{
			e.ID,
			e.Name,
			e.Email,
			baseUrl,
			e.ShortDescription,
		}

		_, _ = fmt.Fprintln(wr, strings.Join(line, "\t"))
	}

	return wr.Flush()
}

func (l ExtensionList) RenderJSON() error {
	data, err := json.Marshal(l.data)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(l.writer, string(data))
	return err
}

func (l ExtensionList) RenderPrettyJSON() error {
	data, err := json.MarshalIndent(l.data, "", " ")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(l.writer, string(data))
	return err
}

func (l ExtensionList) RenderYAML() error {
	data, err := yaml.Marshal(l.data)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(l.writer, string(data))
	return err
}
