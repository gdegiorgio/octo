package install

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gdegiorgio/octo/internal/pkg"
	"github.com/schollz/progressbar/v3"
)

func downloadPackage(p *pkg.Package, platform *pkg.Platform) error {

	packagePath, err := createPackageFolder(p.Name, p.Version)

	if err != nil {
		return fmt.Errorf("could not create package dir : %v", err)
	}

	downloadUrl := platform.URL
	filename := path.Base(downloadUrl)
	filePath := filepath.Join(packagePath, filename)

	if err != nil {
		return fmt.Errorf("invalid download url: %s", downloadUrl)
	}

	req, err := http.NewRequest("GET", downloadUrl, nil)

	if err != nil {
		return fmt.Errorf("could not make http request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return fmt.Errorf("could not download package: %v", err)
	}

	defer resp.Body.Close()

	f, _ := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		fmt.Sprintf("Downloading %s@%s", p.Name, p.Version),
	)

	io.Copy(io.MultiWriter(f, bar), resp.Body)

	if platform.Extract.Format == "zip" {
		err = unzipFile(filePath)
	} else if platform.Extract.Format == "tar" {
		err = untarFile(filePath)
	}
	if err != nil {
		return fmt.Errorf("could not extract package %s: %v", filePath, err)
	}

	if err != nil {
		return fmt.Errorf("could not extract package %s: %v", filePath, err)
	}

	// TODO  Copy in $OCTO_HOME/bin

	return nil
}

func runPreInstallHook() {

}

func runPostInstallHook() {}

func rollback() {}
