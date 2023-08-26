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

func GetProjectSourceRootDir() string {
	var projectSourceDir string

	if Exists(ConfigFilePath) {
		var file, err = os.ReadFile(ConfigFilePath)
		if err != nil {
			panic(err)
		}

		projectSourceDir = string(file)
	} else {
		projectSourceDir = DefaultProjectSourceDir
	}

	return projectSourceDir
}

func GetProjectSourceDir() []string {
	var projectSourceDir []string

	projectSourceRootDir := GetProjectSourceRootDir()

	projectSourceDir = append(projectSourceDir, projectSourceRootDir)

	dirList, err := os.ReadDir(projectSourceRootDir)
	if err != nil {
		panic(err)
	}

	for _, dir := range dirList {
		if dir.IsDir() {
			projectSourceDir = append(projectSourceDir, filepath.Join(projectSourceRootDir, dir.Name()))
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
