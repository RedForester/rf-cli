package version

import (
	"fmt"
	"github.com/redforester/rf-cli/internal/build"
	"github.com/spf13/cobra"
)

func NewCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("version %s (%s)\n", build.Version, build.Date)
		},
	}

	return cmd
}
