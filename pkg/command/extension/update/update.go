package update

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
	Id     string
}

func NewCmdExtUpdate(f *factory.Factory) *cobra.Command {
	var opt Options

	cmd := &cobra.Command{
		Use:   "update [--id=uuid] [--format=json]",
		Short: "Update extension info from manifest",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			run(f, cmd, args, opt)
		},
	}

	cmd.Flags().StringVar(&opt.Format, "format", "pretty", "output format (json, pretty-json, pretty)")
	cmd.Flags().StringVarP(&opt.File, "file", "f", "manifest.yaml", "manifest file path")
	cmd.Flags().StringVar(&opt.File, "id", "", "overwrite extension id from manifest")

	return cmd
}

func run(f *factory.Factory, cmd *cobra.Command, args []string, opt Options) {
	ext, err := manifest.ReadFromFile(opt.File)
	if err != nil {
		fmt.Printf("mainfest error: %s \n", err)
		os.Exit(1)
	}

	if ext.ID == "" && opt.Id != "" {
		ext.ID = opt.Id
	}
	if ext.ID == "" {
		fmt.Println("invalid extension ID")
		os.Exit(1)
	}

	result, err := update(f, ext)
	if err != nil {
		fmt.Printf("api error: %s \n", err)
		os.Exit(1)
	}

	switch opt.Format {
	case "json":
		str, _ := json.Marshal(result)
		fmt.Println(string(str))
	case "pretty-json":
		str, _ := json.MarshalIndent(result, "", " ")
		fmt.Println(string(str))
	case "pretty":
		manifest.PrettyPrint(result)
	default:
		fmt.Println("format not found")
		os.Exit(1)
	}
}

func update(f *factory.Factory, ext *extension.Extension) (*extension.Extension, error) {
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

	result, err := api.Ext.Update(ext)
	if err != nil {
		return nil, err
	}

	return result, nil
}
