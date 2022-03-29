package view

import (
	"fmt"
	"github.com/deissh/rf-cli/pkg/rf"
	"io"
	"strings"
	"text/tabwriter"
)

type ExtensionList struct {
	Data   *[]rf.Extension
	Writer io.Writer
}

func (l ExtensionList) header() []string {
	return []string{
		"ID",
		"NAME",
		"SUPPORT",
		"BASE URL",
		"DESCRIPTION",
	}
}

func (l ExtensionList) printHeader(wr io.Writer) {
	head := strings.Join(l.header(), "\t")
	_, _ = fmt.Fprintln(wr, head)
}

func (l ExtensionList) Render() error {
	wr := tabwriter.NewWriter(l.Writer, 0, 0, 2, ' ', 0)

	l.printHeader(wr)

	for _, e := range *l.Data {
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
