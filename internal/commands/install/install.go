package install

import (
	"github.com/spf13/cobra"
)

func NewInstallCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "install [package]",
		Short: "Install a new package",
		Long:  "Install a new package",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.PrintErr("install command is currently not implemented.")
		},
	}

}
