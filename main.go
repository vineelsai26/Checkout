package main

import (
	"fmt"
	"os"
	"path/filepath"

	"vineelsai.com/checkout/project"
	"vineelsai.com/checkout/utils"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		panic("not enough arguments")
	}

	switch args[0] {
	case "init":
		if len(args) == 1 {
			projectDir := project.PromptForName()
			if utils.Exists(projectDir) {
				fmt.Println("Found Project at:", projectDir)
				fmt.Println("Checking Out Project...")
				project.Init(projectDir)
			} else {
				fmt.Println("Project not found")
			}
			return
		}
		for _, projectSourceRootDir := range utils.GetProjectSourceDir() {
			if utils.Exists(projectSourceRootDir) {
				projectDir := filepath.Join(projectSourceRootDir, args[1])
				if utils.Exists(projectDir) {
					fmt.Println("Found Project at:", projectDir)
					fmt.Println("Checking Out Project...")
					project.Init(projectDir)
				} else {
					fmt.Println("Project not found")
				}
			} else {
				fmt.Println("Project source directory not found")
			}
		}

	case "deinit":
		if len(args) == 2 {
			project.DeInit(args[1], utils.DefaultProjectSource)
		} else {
			project.DeInit(args[1], args[2])
		}
	}
}
