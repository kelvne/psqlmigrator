package fns

import (
	"os"
	"path/filepath"
)

func migrationsPath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	path := filepath.Join(wd, "migrations")

	_, err = os.ReadDir(path)
	if os.IsNotExist(err) {
		os.Mkdir(path, 0755)
		err = nil
	}

	return path, err
}
