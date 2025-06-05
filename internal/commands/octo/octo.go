package octo

import (
	"fmt"

	"github.com/gdegiorgio/octo/internal/commands/install"
	"github.com/gdegiorgio/octo/internal/commands/list"
	"github.com/gdegiorgio/octo/internal/commands/version"
	"github.com/spf13/cobra"
)

func Main() {
	err := newRootCmd().Execute()
	if err != nil {
		fmt.Println("Could not launch octo : %w", err)
	}
}

func newRootCmd() *cobra.Command {

	octo := &cobra.Command{
		Use:   "octo <command> <subcommand> [flags]",
		Short: "Octo CLI",
		Long:  `🐙 Install your packages everywhere`,
	}

	octo.AddCommand(version.NewVersionCmd())
	octo.AddCommand(install.NewInstallCmd())
	octo.AddCommand(list.NewListCommand())

	return octo
}
