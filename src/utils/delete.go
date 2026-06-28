package utils

import (
	"fmt"
	"os"
)

func DeleteFolder(path string) error {
	return DeleteFolderWithOptions(path, false)
}

func DeleteFolderWithOptions(path string, dryRun bool) error {
	if dryRun {
		fmt.Println("Would delete", path)
		return nil
	}

	fmt.Println("Deleting", path)
	return os.RemoveAll(path)
}
