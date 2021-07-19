package list

import (
	"encoding/json"
	"fmt"
	"github.com/deissh/rf-cli/internal/config"
	"github.com/deissh/rf-cli/pkg/extension"
	"github.com/deissh/rf-cli/pkg/factory"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
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

	resp, err := client.Get(cfg.Rf.BaseURL + "/api/extensions")
	if err != nil {
		fmt.Printf("rf err: %s \n", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var data []extension.Extension
		json.Unmarshal(bodyBytes, &data)

		fmt.Println("ID                                     NAME")
		for _, ext := range data {
			fmt.Printf("%s   %s\n", ext.ID, ext.Name)
		}
	} else {
		fmt.Println(resp.Status)
	}
}
