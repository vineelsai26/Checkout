package utils

import "os"

func DeleteFolder(path string) error {
	return os.RemoveAll(path)
}
