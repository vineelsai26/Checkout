//go:build unix

package utils

import (
	"fmt"
	"os"
	"syscall"
)

func preserveOwnership(sourcePath, destPath string, fileInfo os.FileInfo) error {
	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		return fmt.Errorf("failed to get raw syscall.Stat_t data for %q", sourcePath)
	}
	return os.Lchown(destPath, int(stat.Uid), int(stat.Gid))
}
