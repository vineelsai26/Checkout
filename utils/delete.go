package utils

import (
	"fmt"
	"os"
)

func DeleteFolder(path string) error {
	fmt.Println("Deleting", path)
	return os.RemoveAll(path)
}
