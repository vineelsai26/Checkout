package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetHome() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return home
}

func GetProjectDestDir(projectDir string) string {
	return filepath.Join(ProjectCheckoutDir, projectDir)
}

func GetProjectSourceDir() []string {

	var projectDirs []string

	dirList, err := os.ReadDir(ProjectSourceDir)
	if err != nil {
		panic(err)
	}

	for _, dir := range dirList {
		if dir.IsDir() {
			projectDirs = append(projectDirs, filepath.Join(ProjectSourceDir, dir.Name()))
		}
	}
	return projectDirs
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
