//go:build windows

package utils

import "os"

func preserveOwnership(_ string, _ string, _ os.FileInfo) error {
	return nil
}
