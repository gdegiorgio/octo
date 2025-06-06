package update

import (
	"github.com/spf13/cobra"
)

func NewUpdateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "update [package]",
		Short: "Update a package to latest version",
		Long:  "Update a package to latest version",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.PrintErr("update command is currently not implemented.")
		},
	}

}
