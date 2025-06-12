package install

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gdegiorgio/octo/internal/pkg"
	"gopkg.in/yaml.v3"
)

func fetchPackageMetadata(url string) (*pkg.Package, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("package not found (HTTP %d)", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var pkg pkg.Package
	if err := yaml.Unmarshal(body, &pkg); err != nil {
		return nil, err
	}

	return &pkg, nil
}

func retrievePlatformMetadata(pkg *pkg.Package, os string, arch string) *pkg.Platform {
	for _, platform := range pkg.Platforms {
		if (platform.Arch == arch) && (platform.OS == os) {
			return &platform
		}
	}
	return nil
}

func createPackageFolder(packageName string) error {
	homeDir := os.Getenv("OCTO_HOME")
	_, err := os.Stat(fmt.Sprintf("%s/packages/%s", homeDir, packageName))
	if err != nil {
		return os.Mkdir(fmt.Sprintf("%s/packages/%s", homeDir, packageName), 0777)
	}
	return nil
}
