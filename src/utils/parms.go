package utils

import (
	"os"
	"path/filepath"
	"strings"
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

	if envSourceDir := os.Getenv("CHECKOUT_SOURCE_DIR"); envSourceDir != "" {
		projectSourceDir = envSourceDir
	} else if Exists(ConfigFilePath) {
		var file, err = os.ReadFile(ConfigFilePath)
		if err != nil {
			panic(err)
		}

		projectSourceDir = strings.TrimSpace(string(file))
	} else {
		projectSourceDir = DefaultProjectSourceDir
	}

	return filepath.Clean(projectSourceDir)
}()

var ProjectCheckoutRootDir = func() string {
	if checkoutRoot := os.Getenv("CHECKOUT_ROOT"); checkoutRoot != "" {
		return filepath.Clean(checkoutRoot)
	}

	return filepath.Join(HOME, "Personal")
}()
