package project

import (
	"path/filepath"

	"vineelsai.com/checkout/src/utils"
)

func DeInit(projectName string, projectFolder string) {
	projectCheckoutPath := filepath.Join(utils.ProjectCheckoutRootDir, projectName)
	projectSourceDir := filepath.Join(utils.ProjectSourceDir, projectFolder, projectName)

	if err := utils.CopyDirectory(projectCheckoutPath, projectSourceDir); err != nil {
		panic(err)
	}

	if err := utils.DeleteFolder(projectCheckoutPath); err != nil {
		panic(err)
	}
}
