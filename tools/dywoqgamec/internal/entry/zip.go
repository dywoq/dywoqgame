package entry

import (
	"archive/zip"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/dywoq/dywoqgame/tools/dywoqgamec/internal/config"
)

func Zip(c *config.Config) error {
	scripts, err := getScripts(c.Scripts)
	if err != nil {
		return err
	}

	assets, err := getAssets(c.Assets)
	if err != nil {
		return err
	}

	writer, err := createZip("game.zip")
	if err != nil {
		return err
	}
	defer writer.Close()

	err = filesToZip(scripts, writer)
	if err != nil {
		return err
	}

	err = filesToZip(assets, writer)
	if err != nil {
		return err
	}

	return nil
}

func getScripts(folder string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(folder, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return []string{}, err
	}
	return files, err
}

func getAssets(folder string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(folder, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return []string{}, err
	}
	return files, nil
}

func createZip(name string) (*zip.Writer, error) {
	zipFile, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	defer zipFile.Close()
	zipWriter := zip.NewWriter(zipFile)
	return zipWriter, nil
}

func filesToZip(files []string, w *zip.Writer) error {
	if w == nil {
		return errors.New("github.com/dywoq/dywoqgame/tools/dywoqgamec/internal/entry: w is nil")
	}
	
	return nil
}
