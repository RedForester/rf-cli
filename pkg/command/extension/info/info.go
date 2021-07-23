package info

import (
	"encoding/json"
	"fmt"
	"github.com/deissh/rf-cli/pkg/extension"
	"github.com/deissh/rf-cli/pkg/factory"
	"github.com/deissh/rf-cli/pkg/rf_api"
	"github.com/spf13/cobra"
	"os"
)

type Options struct {
	Format string
}

func NewCmdExtInfo(f *factory.Factory) *cobra.Command {
	var opt Options

	cmd := &cobra.Command{
		Use:   "info <id> [--format=json]",
		Short: "Return extension info",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			run(f, cmd, args, opt)
		},
	}

	cmd.Flags().StringVar(&opt.Format, "format", "pretty", "output format (json, pretty-json, pretty)")

	return cmd
}

func run(f *factory.Factory, cmd *cobra.Command, args []string, opt Options) {
	extId := args[0]

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

	data, err := api.Ext.Get(extId)
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

func prettyPrint(ext *extension.Extension) {
	fmt.Println("ID:", ext.ID)
	fmt.Println("Name:", ext.Name)
	fmt.Println("Description:", ext.Description)
	fmt.Println("Email:", ext.Email)
	fmt.Println("BaseURL:", ext.BaseURL)
	fmt.Println("Extension user:")
	fmt.Println("    Username:", ext.User.Username)

	fmt.Println("\nCommands:")
	for i, command := range ext.Commands {
		fmt.Printf(" %d. %s\n", i+1, command.Name)
		fmt.Println("    Description:", command.Description)
		fmt.Println("    Group:", command.Group)
		fmt.Printf("    Type: %+v\n", command.Type)

		fmt.Println("\n    Rules")
		for j, rule := range command.Rules {
			fmt.Printf("     %d. %+v\n", j+1, rule)
		}
		fmt.Println()
	}

	fmt.Println("\nRequired Types:")
	for i, reqType := range ext.RequiredTypes {
		fmt.Printf(" %d. %s\n", i+1, reqType.Name)

		fmt.Println("    Properties")
		for j, prop := range reqType.Properties {
			fmt.Printf("     %d. %s (%s / %s)\n", j+1, prop.Name, prop.Category, prop.Argument)
		}
		fmt.Println()
	}
}
