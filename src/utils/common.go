package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetProjectSourceDir() []string {
	projectSourceDir := make([]string, 0)
	projectSourceDir = append(projectSourceDir, ProjectSourceDir)

	dirList, err := os.ReadDir(ProjectSourceDir)
	if err != nil {
		panic(err)
	}

	for _, dir := range dirList {
		if dir.IsDir() {
			projectSourceDir = append(projectSourceDir, filepath.Join(ProjectSourceDir, dir.Name()))
		}
	}
	return projectSourceDir
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
