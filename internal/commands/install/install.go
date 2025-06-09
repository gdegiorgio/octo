package install

import (
	"fmt"
	"runtime"

	"github.com/gdegiorgio/octo/pkg/logger"
	"github.com/spf13/cobra"
)

func NewInstallCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "install [package]",
		Short: "Install a new package",
		Long:  "Install a new package",
		Args:  cobra.MinimumNArgs(1),
		Run: RunInstall,
	}

}


func RunInstall(cmd *cobra.Command, args []string) {

	goos := runtime.GOOS
	goarch := runtime.GOARCH


	logger.Debug(cmd, fmt.Sprintf("Detected os/platform: %s %s", goos, goarch))

}
