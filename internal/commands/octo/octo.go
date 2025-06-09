package octo

import (
	"fmt"

	"github.com/gdegiorgio/octo/internal/commands/install"
	"github.com/gdegiorgio/octo/internal/commands/list"
	"github.com/gdegiorgio/octo/internal/commands/update"
	"github.com/gdegiorgio/octo/internal/commands/upgrade"
	"github.com/gdegiorgio/octo/internal/commands/version"
	"github.com/gdegiorgio/octo/pkg/logger"
	"github.com/spf13/cobra"
)

var (
	Verbose bool
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
	octo.AddCommand(update.NewUpdateCmd())
	octo.AddCommand(upgrade.NewUpgradeCmd())


	octo.PersistentFlags().BoolVarP(&logger.Verbose, "verbose", "v", false, "Enable verbose logging")


	return octo
}
