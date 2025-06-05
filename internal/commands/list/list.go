package list

import "github.com/spf13/cobra"

func NewListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all installed packages",
		Long:  "List all installed packages",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.PrintErr("list command is currently not implemented.")
		},
	}
}
