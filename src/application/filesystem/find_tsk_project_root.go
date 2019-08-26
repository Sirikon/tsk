package filesystem

import (
	"errors"
	"os"
	"path"
	"path/filepath"
)

func FindTskProjectRoot(basePath string) (string, error) {
	rootFolder, err := getSystemRootFolder()
	if err != nil {
		return "", err
	}

	currentTskFilePath := path.Join(basePath, "Tskfile.yml")

	stat, err := os.Stat(currentTskFilePath)
	isDir := stat != nil && stat.IsDir()

	if os.IsNotExist(err) || isDir {
		if rootFolder == basePath {
			return findTskProjectError("Tskfile.yml not found")
		}
		return FindTskProjectRoot(path.Join(basePath, ".."))
	}

	if err != nil {
		return findTskProjectError("unexpected error while finding Tskfile.yml: " + err.Error())
	}

	return basePath, nil
}

func getSystemRootFolder() (string, error) {
	return filepath.Abs("/")
}

func findTskProjectError(message string) (string, error) {
	return "", errors.New(message)
}
