package remove

import (
	"github.com/spf13/cobra"
)

func NewRemoveCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "remove [package]",
		Short: "Uninstall a package",
		Long:  "Uninstall a package",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.PrintErr("remove command is currently not implemented.")
		},
	}
}
