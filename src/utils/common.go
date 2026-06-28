package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetProjectSourceDir() ([]string, error) {
	projectSourceDir := make([]string, 0)
	projectSourceDir = append(projectSourceDir, ProjectSourceDir)

	if !Exists(ProjectSourceDir) {
		return nil, fmt.Errorf("project source directory not found: %s", ProjectSourceDir)
	}

	dirList, err := os.ReadDir(ProjectSourceDir)
	if err != nil {
		return nil, err
	}

	for _, dir := range dirList {
		if dir.IsDir() {
			projectSourceDir = append(projectSourceDir, filepath.Join(ProjectSourceDir, dir.Name()))
		}
	}
	return projectSourceDir, nil
}

func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func CreateIfNotExists(dir string, perm os.FileMode) error {
	if Exists(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}
