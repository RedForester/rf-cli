package extension

import (
	"fmt"
	"github.com/deissh/rf-cli/internal/config"
	"github.com/deissh/rf-cli/internal/factory"
	"github.com/deissh/rf-cli/internal/utils"
	"github.com/deissh/rf-cli/pkg/rf/extension"
	"github.com/deissh/rf-cli/pkg/view"
	"github.com/spf13/cobra"
)

func NewCmdExtension() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "extension",
		Short:   "Manage extension",
		Aliases: []string{"ext", "e"},
		Run:     run,
	}

	cmd.Flags().StringP("format", "f", "pretty", "output format (json, pretty-json, yaml, pretty)")

	return cmd
}

func run(cmd *cobra.Command, _ []string) {
	client := factory.NewClient(config.Config.Client.Username, config.Config.Client.PasswordHash)

	data, err := func() (*[]extension.Extension, error) {
		s := utils.PrintSpinner("Fetching extensions...")
		defer s.Stop()

		return client.Ext.GetAll()
	}()
	utils.ExitIfError(err)

	r := view.NewExtensionList(data)

	format, _ := cmd.Flags().GetString("format")
	switch format {
	case "json":
		err = r.RenderJSON()
		utils.ExitIfError(err)
		return
	case "pretty-json":
		err = r.RenderPrettyJSON()
		utils.ExitIfError(err)
		return
	case "yaml":
		err = r.RenderYAML()
		utils.ExitIfError(err)
		return
	case "pretty":
		err = r.Render()
		utils.ExitIfError(err)
	default:
		fmt.Println("Invalid output format")
	}
}
