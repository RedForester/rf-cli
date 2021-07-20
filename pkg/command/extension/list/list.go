package list

import (
	"fmt"
	"github.com/deissh/rf-cli/internal/config"
	"github.com/deissh/rf-cli/pkg/extension"
	"github.com/deissh/rf-cli/pkg/factory"
	"github.com/deissh/rf-cli/pkg/rf_api"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

type Options struct {
	HttpClient func() (*http.Client, error)
	Config     func() (config.Config, error)

	Limit int
	Owner string
}

func NewCmdExtList(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Return all registered extensions",
		Run: func(cmd *cobra.Command, args []string) {
			run(f, cmd, args)
		},
	}

	return cmd
}

func run(f *factory.Factory, cmd *cobra.Command, args []string) {
	cfg, err := f.Config()
	if err != nil {
		fmt.Printf("internal error, err: %s \n", err)
		os.Exit(1)
	}

	client, err := f.HttpClient()
	if err != nil {
		fmt.Printf("internal error, err: %s \n", err)
		os.Exit(1)
	}

	opt := rf_api.NewOptions(cfg.Rf.BaseURL)
	api := rf_api.New(client, opt)

	data, err := api.Ext.GetAll()
	if err != nil {
		fmt.Printf("internal error, err: %s \n", err)
		os.Exit(1)
	}

	printExtensions(data)
}

func printExtensions(data *[]extension.Extension) {
	fmt.Println("ID                                     NAME")
	for _, ext := range *data {
		fmt.Printf("%s   %s\n", ext.ID, ext.Name)
	}
}
