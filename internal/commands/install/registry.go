package install

import (
	"context"
	"fmt"
	"io"
	"net/http"
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
