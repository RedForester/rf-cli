package version

import (
	"fmt"
	"github.com/deissh/rf-cli/internal/build"
	"github.com/deissh/rf-cli/internal/factory"
	"github.com/spf13/cobra"
)

func NewCmdVersion(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("version %s (%s)\n", build.Version, build.Date)
		},
	}

	return cmd
}
