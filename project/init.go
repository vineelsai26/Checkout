package project

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
	"vineelsai.com/checkout/utils"
)

func getProjectsList() []string {
	var projects []string

	for _, projectSourceRootDir := range utils.GetProjectSourceDir() {
		if utils.Exists(projectSourceRootDir) {
			files, err := os.ReadDir(projectSourceRootDir)
			if err != nil {
				panic(err)
			}

			projects = append(projects, projectSourceRootDir)

			for _, file := range files {
				if file.IsDir() {
					projects = append(projects, filepath.Join(projectSourceRootDir, file.Name()))
				}
			}
		} else {
			fmt.Println("Project source directory not found")
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

func Init(projectDir string) {
	projectCheckoutDir := filepath.Join(utils.ProjectCheckoutRootDir, strings.Split(projectDir, "/")[len(strings.Split(projectDir, "/"))-1])

	if utils.Exists(projectDir) {
		err := utils.CopyDirectory(projectDir, projectCheckoutDir)
		if err != nil {
			panic(err)
		}

		if err := utils.DeleteFolder(projectDir); err != nil {
			panic(err)
		}

		exec.Command("code", projectCheckoutDir).Run()
		return
	}

	fmt.Println("Project not found")
}
