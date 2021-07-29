package update

import (
	"encoding/json"
	"fmt"
	"github.com/deissh/rf-cli/pkg/extension"
	"github.com/deissh/rf-cli/pkg/factory"
	"github.com/deissh/rf-cli/pkg/rf_api"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
	ext, err := readManifest(opt.File)
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

	if err := update(f, ext); err != nil {
		fmt.Printf("api error: %s \n", err)
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
		extension.PrettyPrint(ext)
	default:
		fmt.Println("format not found")
		os.Exit(1)
	}
}

func readManifest(file string) (*extension.Extension, error) {
	var data extension.Extension

	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		return nil, err
	}

	return &data, err
}

func update(f *factory.Factory, ext *extension.Extension) error {
	cfg, err := f.Config()
	if err != nil {
		return err
	}

	client, err := f.HttpClient()
	if err != nil {
		return err
	}

	apiOpts := rf_api.NewOptions(cfg.Rf.BaseURL)
	api := rf_api.New(client, apiOpts)

	ext, err = api.Ext.Update(ext)
	if err != nil {
		return err
	}

	return nil
}
