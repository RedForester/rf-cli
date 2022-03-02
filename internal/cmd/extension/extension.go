package extension

import (
	"fmt"
	"github.com/deissh/rf-cli/internal/config"
	"github.com/deissh/rf-cli/internal/factory"
	"github.com/deissh/rf-cli/internal/utils"
	"github.com/deissh/rf-cli/pkg/rf/extension"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"text/tabwriter"
)

func NewCmdExtension() *cobra.Command {
	return &cobra.Command{
		Use:     "extension",
		Short:   "Manage extension",
		Aliases: []string{"ext", "e"},
		Run:     run,
	}
}

func run(cmd *cobra.Command, _ []string) {
	client := factory.NewClient(config.Config.Client.Username, config.Config.Client.PasswordHash)

	data, err := func() (*[]extension.Extension, error) {
		s := utils.PrintSpinner("Fetching extensions...")
		defer s.Stop()

		return client.Ext.GetAll()
	}()
	utils.ExitIfError(err)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tEMAIL\tBASE URL")
	for _, extension := range *data {
		baseUrl := ""
		if extension.BaseURL != nil {
			baseUrl = *extension.BaseURL
		}
		line := []string{
			extension.ID,
			extension.Name,
			extension.Email,
			baseUrl,
		}
		fmt.Fprintln(w, strings.Join(line, "\t"))
	}

	w.Flush()
}
