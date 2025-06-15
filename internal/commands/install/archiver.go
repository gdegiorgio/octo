package install

import (
	"archive/tar"
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/schollz/progressbar/v3"
)

func unzipFile(fp string) error {

	r, err := zip.OpenReader(fp)
	if err != nil {
		return err
	}
	defer r.Close()

	bar := progressbar.DefaultBytes(
		int64(len(r.File)),
	)

	for _, f := range r.File {

		bar.Add(1)
		bar.Describe(fmt.Sprintf("Extracting %s", f.Name))

		rc, err := f.Open()

		if err != nil {
			return fmt.Errorf("cannot extract files: %v", err)
		}

		defer rc.Close()

		newFilePath := filepath.Join(filepath.Dir(fp), f.Name)

		if f.FileInfo().IsDir() {
			err = os.MkdirAll(newFilePath, 0777)
			if err != nil {
				return fmt.Errorf("cannot extract internal dirs: %v", err)
			}
		} else {
			uncompressedFile, err := os.Create(newFilePath)
			if err != nil {
				return fmt.Errorf("cannot extract files: %v", err)
			}
			_, err = io.Copy(uncompressedFile, rc)
			if err != nil {
				return fmt.Errorf("cannot extract files: %v", err)
			}
		}

	}

	return nil
}

func untarFile(fp string) error {

	f, err := os.Open(fp)
	if err != nil {
		return fmt.Errorf("could not open tar file: %v", err)
	}
	defer f.Close()

	r := tar.NewReader(f)
	if err != nil {
		return err
	}

	for {
		header, err := r.Next()

		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("could not read tar file: %v", err)
		}

		bar := progressbar.DefaultBytes(
			header.Size,
			fmt.Sprintf("Extracting %s", header.Name),
		)

		newFilePath := filepath.Join(filepath.Dir(fp), header.Name)

		if header.Typeflag == tar.TypeDir {
			err = os.MkdirAll(newFilePath, 0777)
			if err != nil {
				return fmt.Errorf("cannot extract internal dirs: %v", err)
			}
			continue
		}

		uncompressedFile, err := os.Create(newFilePath)
		if err != nil {
			return fmt.Errorf("cannot extract files: %v", err)
		}

		n, err := io.Copy(uncompressedFile, r)
		if err != nil {
			return fmt.Errorf("cannot extract files: %v", err)
		}

		bar.Add(int(n))
		uncompressedFile.Close()
	}

	return nil
}
