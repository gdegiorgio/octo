package version

import (
	"github.com/spf13/cobra"
)

func NewVersionCmd(version string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show current Octo CLI version",
		Long:  "Show current Octo CLI version",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("🐙", version)
		},
	}
}
