package octo

import (
	"fmt"

	"github.com/gdegiorgio/octo/internal/commands/install"
	"github.com/gdegiorgio/octo/internal/commands/list"
	"github.com/gdegiorgio/octo/internal/commands/version"
	"github.com/spf13/cobra"
)

func Main(octoVersion string) {
	err := newRootCmd(octoVersion).Execute()
	if err != nil {
		fmt.Println("Could not launch octo : %w", err)
	}
}

func newRootCmd(octoVersion string) *cobra.Command {

	octo := &cobra.Command{
		Use:   "octo <command> <subcommand> [flags]",
		Short: "Octo CLI",
		Long:  `🐙 Install your packages everywhere`,
	}

	octo.AddCommand(version.NewVersionCmd(octoVersion))
	octo.AddCommand(install.NewInstallCmd())
	octo.AddCommand(list.NewListCommand())

	return octo
}
