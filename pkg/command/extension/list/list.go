package list

import (
	"encoding/json"
	"fmt"
	"github.com/deissh/rf-cli/internal/config"
	"github.com/deissh/rf-cli/pkg/extension"
	"github.com/deissh/rf-cli/pkg/factory"
	"github.com/deissh/rf-cli/pkg/rf_api"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

type Options struct {
	HttpClient func() (*http.Client, error)
	Config     func() (config.Config, error)

	Limit  int
	Owned  bool
	Format string
}

func NewCmdExtList(f *factory.Factory) *cobra.Command {
	var opt Options

	cmd := &cobra.Command{
		Use:   "list [--format=json]",
		Short: "Return all registered extension",
		Run: func(cmd *cobra.Command, args []string) {
			run(f, cmd, args, opt)
		},
	}

	cmd.Flags().StringVar(&opt.Format, "format", "pretty", "output format (json, pretty-json, pretty)")
	cmd.Flags().BoolVar(&opt.Owned, "owned", false, "show only current user extension")

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

	var data *[]extension.Extension

	if opt.Owned {
		data, err = api.Ext.GetOwned()
	} else {
		data, err = api.Ext.GetAll()
	}

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
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"ID", "NAME", "AUTHOR", "BASE URL"})
	for _, ext := range *data {
		baseUrl := "NOT SET"
		if ext.BaseURL != nil {
			baseUrl = *ext.BaseURL
		}

		table.Append([]string{
			ext.ID,
			ext.Name,
			ext.Email,
			baseUrl,
		})
	}

	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(3)
	table.SetAlignment(3)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("  ")
	table.SetNoWhiteSpace(true)
	table.Render()
}
