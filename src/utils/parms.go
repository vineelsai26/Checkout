package utils

import (
	"os"
	"path/filepath"
)

var DefaultProjectSource = "vineelsai26"

var DefaultProjectSourceDir = "/Users/vineel/Dropbox/GitHub"

var HOME = func() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return home
}()

var ConfigFilePath = filepath.Join(HOME, ".checkout", "source_dir")

var ProjectSourceDir = func() string {
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
}()

var ProjectCheckoutRootDir = filepath.Join(HOME, "Personal")
