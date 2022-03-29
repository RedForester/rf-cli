package view

import (
	"errors"
	"github.com/deissh/rf-cli/internal/factory"
	"github.com/deissh/rf-cli/internal/utils"
	"github.com/deissh/rf-cli/pkg/rf"
	"github.com/deissh/rf-cli/pkg/view"
	"github.com/spf13/cobra"
	"os"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "view {extId} [-f format]",
		Short:   "View information about extension",
		Aliases: []string{"show"},
		Args:    cobra.ExactArgs(1),
		Run:     run,
	}

	cmd.Flags().StringP("format", "f", "pretty", "output format (json, pretty-json, yaml, pretty)")

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	client := factory.ClientInstance

	extId := args[0]

	data, err := func() (*rf.Extension, error) {
		s := utils.PrintSpinner("Fetching extension information from RedForester...")
		defer s.Stop()

		return client.Ext.Get(extId)
	}()
	utils.ExitIfError(err)

	format, _ := cmd.Flags().GetString("format")

	err = render(format, view.Extension{
		Data:   data,
		Writer: os.Stdout,
	})
	utils.ExitIfError(err)
}

func render(format string, r view.Extension) error {
	switch format {
	case "json":
		return view.RenderJSON(r.Writer, r.Data)
	case "pretty-json":
		return view.RenderPrettyJSON(r.Writer, r.Data)
	case "yaml":
		return view.RenderYAML(r.Writer, r.Data)
	case "pretty":
		return r.Render()
	default:
		return errors.New("invalid output format")
	}
}
