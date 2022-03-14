package list

import (
	"errors"
	"fmt"
	"github.com/deissh/rf-cli/internal/factory"
	"github.com/deissh/rf-cli/internal/utils"
	"github.com/deissh/rf-cli/pkg/rf"
	"github.com/deissh/rf-cli/pkg/view"
	"github.com/spf13/cobra"
	"os"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Show all available extension",
		Aliases: []string{"l"},
		Run:     run,
	}

	cmd.Flags().Bool("owned", false, "return all owned extensions")
	cmd.Flags().StringP("format", "f", "pretty", "output format (json, pretty-json, yaml, pretty)")

	return cmd
}

func run(cmd *cobra.Command, _ []string) {
	client := factory.ClientInstance

	owned, _ := cmd.Flags().GetBool("owned")

	data, err := func() (*[]rf.Extension, error) {
		s := utils.PrintSpinner("Fetching extensions...")
		defer s.Stop()

		if owned {
			return client.Ext.GetOwned()
		}

		return client.Ext.GetAll()
	}()
	utils.ExitIfError(err)

	if len(*data) == 0 {
		fmt.Println()
		utils.Exit("No result found")
		return
	}

	format, _ := cmd.Flags().GetString("format")

	err = render(format, view.ExtensionList{
		Data:   data,
		Writer: os.Stdout,
	})

	utils.ExitIfError(err)
}

func render(format string, r view.ExtensionList) error {
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
