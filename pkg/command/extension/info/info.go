package info

import (
	"encoding/json"
	"fmt"
	"github.com/deissh/rf-cli/internal/factory"
	"github.com/deissh/rf-cli/pkg/extension"
	"github.com/deissh/rf-cli/pkg/manifest"
	"github.com/deissh/rf-cli/pkg/rf_api"
	"github.com/spf13/cobra"
	"os"
)

type Options struct {
	Format string
	File   string
}

func NewCmdExtInfo(f *factory.Factory) *cobra.Command {
	var opt Options

	cmd := &cobra.Command{
		Use:   "info [id] [--format=json]",
		Short: "Return extension info from manifest or by id",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			run(f, cmd, args, opt)
		},
	}

	cmd.Flags().StringVar(&opt.Format, "format", "pretty", "output format (json, pretty-json, pretty)")
	cmd.Flags().StringVarP(&opt.File, "file", "f", "manifest.yaml", "manifest file path")

	return cmd
}

func run(f *factory.Factory, cmd *cobra.Command, args []string, opt Options) {
	var ext *extension.Extension
	var err error

	switch len(args) {
	case 0:
		ext, err = manifest.ReadFromFile(opt.File)
	case 1:
		ext, err = fetch(f, args[0])
	}

	if err != nil {
		fmt.Printf("mainfest error: %s \n", err)
		os.Exit(1)
	}

	switch opt.Format {
	case "json":
		str, _ := json.Marshal(ext)
		fmt.Println(string(str))
	case "pretty-json":
		str, _ := json.MarshalIndent(ext, "", " ")
		fmt.Println(string(str))
	case "pretty":
		manifest.PrettyPrint(ext)
	default:
		fmt.Println("format not found")
		os.Exit(1)
	}
}

func fetch(f *factory.Factory, id string) (*extension.Extension, error) {
	cfg, err := f.Config()
	if err != nil {
		return nil, err
	}

	client, err := f.HttpClient()
	if err != nil {
		return nil, err
	}

	apiOpts := rf_api.NewOptions(cfg.Rf.BaseURL)
	api := rf_api.New(client, apiOpts)

	data, err := api.Ext.Get(id)
	if err != nil {
		return nil, err
	}

	return data, err
}
