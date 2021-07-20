package list

import (
	"encoding/json"
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

	Limit  int
	Owner  string
	Format string
}

func NewCmdExtList(f *factory.Factory) *cobra.Command {
	var opt Options

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Return all registered extensions",
		Run: func(cmd *cobra.Command, args []string) {
			run(f, cmd, args, opt)
		},
	}

	cmd.Flags().StringVar(&opt.Format, "format", "pretty", "output format (json, pretty-json, pretty)")

	return cmd
}

func run(f *factory.Factory, cmd *cobra.Command, args []string, opt Options) {
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

	apiOpts := rf_api.NewOptions(cfg.Rf.BaseURL)
	api := rf_api.New(client, apiOpts)

	data, err := api.Ext.GetAll()
	if err != nil {
		fmt.Printf("internal error, err: %s \n", err)
		os.Exit(1)
	}

	switch opt.Format {
	case "json":
		str, _ := json.Marshal(data)
		fmt.Println(string(str))
	case "pretty-json":
		str, _ := json.MarshalIndent(data, "", " ")
		fmt.Println(string(str))
	case "pretty":
		prettyPrint(data)
	default:
		fmt.Println("format not found")
		os.Exit(1)
	}
}

func prettyPrint(data *[]extension.Extension) {
	fmt.Println("ID                                     NAME")
	for _, ext := range *data {
		fmt.Printf("%s   %s\n", ext.ID, ext.Name)
	}
}
