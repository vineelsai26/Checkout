package main

import (
	"fmt"
	"os"
	"path/filepath"

	"vineelsai.com/checkout/utils"
)

func Init(projectName string) {
	for _, projectRootDir := range utils.GetProjectSourceDir() {
		projectDir := filepath.Join(projectRootDir, projectName)

		if utils.Exists(projectDir) {
			err := utils.CopyDirectory(projectDir, utils.GetProjectDestDir(projectName))
			if err != nil {
				panic(err)
			}
			return
		}
	}

	fmt.Println("Project not found")
}

func DeInit(projectName string, projectFolder string) {
	utils.CopyDirectory(utils.ProjectCheckoutDir, filepath.Join(utils.ProjectSourceDir, projectFolder))
}

func main() {
	args := os.Args[1:]

	fmt.Println()

	if len(args) < 2 {
		panic("not enough arguments")
	}

	switch args[0] {
	case "init":
		Init(args[1])
	case "deinit":
		if len(args) == 2 {
			DeInit(args[1], utils.DefaultProjectSource)
		} else {
			DeInit(args[1], args[2])
		}
	}
}
