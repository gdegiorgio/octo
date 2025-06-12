package install

import (
	"fmt"
	"time"

	"github.com/gdegiorgio/octo/internal/pkg"
	"github.com/schollz/progressbar/v3"
)

func downloadPackage(p *pkg.Package) error {
	bar := progressbar.Default(100, fmt.Sprintf("Downloading %s@%s", p.Name, p.Version))
	for _ = range(100) {
    	bar.Add(1)
     	time.Sleep(40 * time.Millisecond)
	}
	return nil
}

func runPreInstallHook() {

}

func runPostInstallHook() {}

func rollback() {}
