package install

import (
	"fmt"
	"runtime"

	"github.com/gdegiorgio/octo/pkg/logger"
	"github.com/spf13/cobra"
)

const defaultRegistryBaseURL = "https://raw.githubusercontent.com/gdegiorgio/octo-pus/main/packages"

var log *logger.Logger

func NewInstallCmd() *cobra.Command {
	installCmd := &cobra.Command{
		Use:   "install [package]",
		Short: "Install a new package",
		Long:  "Install a new package from the octo-pus registry.",
		Example: `
  octo install lazygit
  octo install lazygit@0.24.0
		`,
		Args: cobra.ExactArgs(1),
		Run:  runInstall,
	}

	log = logger.NewLogger(installCmd)

	return installCmd
}

func runInstall(cmd *cobra.Command, args []string) {

	goos := runtime.GOOS
	goarch := runtime.GOARCH

	log.Debug(fmt.Sprintf("Detected platform: %s/%s", goos, goarch))

	packageName, version := getPackageAndVersion(args[0])
	log.Debug(fmt.Sprintf("Installing package: %s (version: %s)", packageName, version))

	octoHome, err := getOctoHome()
	if err != nil {
		log.Error(fmt.Sprintf("Failed to resolve OCTO_HOME: %v", err))
		return
	}

	if isInstalled(octoHome, packageName) {
		log.Info(fmt.Sprintf("Package %s is already installed", packageName))
		return
	}

	url := fmt.Sprintf("%s/%c/%s/%s.yaml", defaultRegistryBaseURL, packageName[0], packageName, version)

	log.Debug(fmt.Sprintf("Looking for %s at %s", packageName, url))

	pkg, err := fetchPackageMetadata(url)

	if err != nil {
		log.Error(fmt.Sprintf("Unable to fetch %s metadata : %v", packageName, err))
		return
	}

	log.Info(fmt.Sprintf("Package %s is available — proceeding with install...", pkg.Name))


	platformMetadata := retrievePlatformMetadata(pkg, goos, goarch)

	if platformMetadata == nil {
		log.Info(fmt.Sprintf("Package %s is not compatible with %s/%s", packageName, goos, goarch))
		return
	}

	if platformMetadata.Hooks.PreInstall != "" {
		log.Info(fmt.Sprintf("Found a pre-install hook."))
		runPreInstallHook()
	}

	err = downloadPackage(pkg)

	if err != nil {
		log.Error(fmt.Sprintf("Failed to download package %s : %v", packageName, err))
		return
	}

	if platformMetadata.Hooks.PostInstall != "" {
		log.Info(fmt.Sprintf("Found a post-install hook."))
		runPostInstallHook()
	}

}
