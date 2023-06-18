package main

import (
	"fmt"
	"os"
	"path/filepath"

	"vineelsai.com/checkout/utils"
)

func initProject(projectName string) {
	for _, projectSourceRootDir := range utils.GetProjectSourceDir() {
		projectSourceDir := filepath.Join(projectSourceRootDir, projectName)
		ProjectCheckoutDir := filepath.Join(utils.ProjectCheckoutRootDir, projectName)

		if utils.Exists(projectSourceDir) {
			err := utils.CopyDirectory(projectSourceDir, ProjectCheckoutDir)
			if err != nil {
				panic(err)
			}

			if err := utils.DeleteFolder(projectSourceDir); err != nil {
				panic(err)
			}
			return
		}
	}

	fmt.Println("Project not found")
}

func deInitProject(projectName string, projectFolder string) {
	projectCheckoutPath := filepath.Join(utils.ProjectCheckoutRootDir, projectName)
	projectSourceDir := filepath.Join(utils.ProjectSourceDir, projectFolder, projectName)

	if err := utils.CopyDirectory(projectCheckoutPath, projectSourceDir); err != nil {
		panic(err)
	}

	if err := utils.DeleteFolder(projectCheckoutPath); err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args[1:]

	fmt.Println()

	if len(args) < 2 {
		panic("not enough arguments")
	}

	switch args[0] {
	case "init":
		initProject(args[1])
	case "deinit":
		if len(args) == 2 {
			deInitProject(args[1], utils.DefaultProjectSource)
		} else {
			deInitProject(args[1], args[2])
		}
	}
}
