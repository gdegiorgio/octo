package install

import (
	"os"
	"path/filepath"
	"strings"
)

func getPackageAndVersion(arg string) (string, string) {
	if strings.Contains(arg, "@") {
		parts := strings.SplitN(arg, "@", 2)
		return parts[0], parts[1]
	}
	return arg, "latest"
}

func getOctoHome() (string, error) {
	if val := os.Getenv("OCTO_HOME"); val != "" {
		return val, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".octo"), nil
}

func isInstalled(octoHome, pkg string) bool {
	path := filepath.Join(octoHome, "bin", pkg)
	_, err := os.Stat(path)
	return err == nil
}
