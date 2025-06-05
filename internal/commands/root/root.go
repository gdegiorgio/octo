package root

import (
	"github.com/gdegiorgio/octo/internal/commands/version"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command{

	root := &cobra.Command{
		Use:   "octo <command> <subcommand> [flags]",
		Short: "Octo CLI",
		Long:  `🐙 Install your packages everywhere`,
	}

	root.AddCommand(version.NewVersionCmd())

	return root
}
