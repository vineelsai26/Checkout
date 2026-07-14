package project

import (
	"fmt"
	"path/filepath"
	"strings"

	"vineelsai.com/checkout/src/utils"
)

func DeInit(projectName string, projectFolder string, options Options) error {
	if projectName == "" || projectName == "." || projectName == ".." || filepath.Base(projectName) != projectName {
		return fmt.Errorf("invalid project name: %q", projectName)
	}
	if filepath.IsAbs(projectFolder) {
		return fmt.Errorf("project folder must be relative: %q", projectFolder)
	}
	cleanFolder := filepath.Clean(projectFolder)
	if cleanFolder == ".." || strings.HasPrefix(cleanFolder, ".."+string(filepath.Separator)) {
		return fmt.Errorf("project folder escapes source root: %q", projectFolder)
	}

	projectCheckoutPath := filepath.Join(utils.ProjectCheckoutRootDir, projectName)
	projectSourceDir := filepath.Join(utils.ProjectSourceDir, cleanFolder, projectName)

	if !utils.Exists(projectCheckoutPath) {
		return fmt.Errorf("checkout project not found: %s", projectCheckoutPath)
	}
	if utils.Exists(projectSourceDir) {
		return fmt.Errorf("source destination already exists: %s", projectSourceDir)
	}

	if err := utils.CopyDirectoryWithOptions(projectCheckoutPath, projectSourceDir, options.Copy); err != nil {
		return err
	}

	return utils.DeleteCopiedTree(projectCheckoutPath, options.Copy)
}
