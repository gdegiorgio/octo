package upgrade

import "github.com/spf13/cobra"

func NewUpgradeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrade Octo CLI to latest version",
		Long:  "Upgrade Octo CLI to latest version",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.PrintErr("upgrade command is currently not implemented.")
		},
	}
}
