package project

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"vineelsai.com/checkout/src/utils"
)

type Options struct {
	Copy   utils.CopyOptions
	OpenVS bool
}

func getProjectsList() ([]string, error) {
	var projects []string

	projectSourceDirs, err := utils.GetProjectSourceDir()
	if err != nil {
		return nil, err
	}

	for _, projectSourceRootDir := range projectSourceDirs {
		if utils.Exists(projectSourceRootDir) {
			files, err := os.ReadDir(projectSourceRootDir)
			if err != nil {
				return nil, err
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

	return projects, nil
}

func PromptForName() (string, error) {
	projects, err := getProjectsList()
	if err != nil {
		return "", err
	}

	prompt := promptui.Select{
		Label: "Select Project",
		Items: projects,
		Size:  20,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

func FindProject(projectName string) (string, error) {
	projectSourceDirs, err := utils.GetProjectSourceDir()
	if err != nil {
		return "", err
	}

	for _, projectSourceRootDir := range projectSourceDirs {
		if !utils.Exists(projectSourceRootDir) {
			continue
		}

		projectDir := filepath.Join(projectSourceRootDir, projectName)
		if utils.Exists(projectDir) {
			return projectDir, nil
		}
	}

	return "", fmt.Errorf("project %q was not found under %s", projectName, utils.ProjectSourceDir)
}

func Init(projectDir string, options Options) error {
	if !utils.Exists(projectDir) {
		return fmt.Errorf("project not found: %s", projectDir)
	}

	projectCheckoutDir := filepath.Join(utils.ProjectCheckoutRootDir, filepath.Base(projectDir))
	if utils.Exists(projectCheckoutDir) {
		return fmt.Errorf("checkout destination already exists: %s", projectCheckoutDir)
	}

	if err := utils.CopyDirectoryWithOptions(projectDir, projectCheckoutDir, options.Copy); err != nil {
		return err
	}

	if err := utils.DeleteCopiedTree(projectDir, options.Copy); err != nil {
		return err
	}

	if options.OpenVS && !options.Copy.DryRun {
		if err := exec.Command("code", projectCheckoutDir).Run(); err != nil {
			return fmt.Errorf("open project in VS Code: %w", err)
		}
	}

	return nil
}
