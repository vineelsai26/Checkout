package project

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"vineelsai.com/checkout/utils"
)

func getProjectsList() []string {
	var projects []string

	for _, projectSourceRootDir := range utils.GetProjectSourceDir() {
		projectSourceRootDir := filepath.Join(projectSourceRootDir)

		if utils.Exists(projectSourceRootDir) {
			files, err := os.ReadDir(projectSourceRootDir)
			if err != nil {
				panic(err)
			}

			for _, file := range files {
				if file.IsDir() {
					projects = append(projects, file.Name())
				}
			}

		}
	}

	return projects
}

func PromptForName() string {
	prompt := promptui.Select{
		Label: "Select Project",
		Items: getProjectsList(),
		Size:  20,
	}

	_, result, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	return result
}

func Init(projectName string) {
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

			exec.Command("code", ProjectCheckoutDir).Run()
			return
		}
	}

	fmt.Println("Project not found")
}
