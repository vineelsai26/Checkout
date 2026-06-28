package project

import (
	"fmt"
	"path/filepath"

	"vineelsai.com/checkout/src/utils"
)

func DeInit(projectName string, projectFolder string, options Options) error {
	projectCheckoutPath := filepath.Join(utils.ProjectCheckoutRootDir, projectName)
	projectSourceDir := filepath.Join(utils.ProjectSourceDir, projectFolder, projectName)

	if !utils.Exists(projectCheckoutPath) {
		return fmt.Errorf("checkout project not found: %s", projectCheckoutPath)
	}
	if utils.Exists(projectSourceDir) {
		return fmt.Errorf("source destination already exists: %s", projectSourceDir)
	}

	if err := utils.CopyDirectoryWithOptions(projectCheckoutPath, projectSourceDir, options.Copy); err != nil {
		return err
	}

	return utils.DeleteFolderWithOptions(projectCheckoutPath, options.Copy.DryRun)
}
