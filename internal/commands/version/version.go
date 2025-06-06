package version

import (
	"github.com/spf13/cobra"
)

// Actual octo version is set by ldflags at build time
var Version = "development"

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show current Octo CLI version",
		Long:  "Show current Octo CLI version",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println(Version)
		},
	}
}
