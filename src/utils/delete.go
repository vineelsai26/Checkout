package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
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

// DeleteCopiedTree removes only entries that CopyDirectoryWithOptions would
// have copied. Skipped files (notably .env files and build caches) remain at
// the source so moving a project can never silently destroy omitted data.
func DeleteCopiedTree(path string, options CopyOptions) error {
	return deleteCopiedTree(path, options, true)
}

func deleteCopiedTree(path string, options CopyOptions, isRoot bool) error {
	// The caller has already selected the project root for copying. Only apply
	// exclusion rules to its descendants; a project directory may itself be
	// named "target", "build", or another excluded cache name.
	if !isRoot && options.ShouldSkip(path) {
		fmt.Println("Preserving skipped path", path)
		return nil
	}

	info, err := os.Lstat(path)
	if err != nil {
		return err
	}

	if !info.IsDir() {
		if options.DryRun {
			fmt.Println("Would delete copied path", path)
			return nil
		}
		return os.Remove(path)
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if err := deleteCopiedTree(filepath.Join(path, entry.Name()), options, false); err != nil {
			return err
		}
	}

	if options.DryRun {
		fmt.Println("Would remove copied directory if empty", path)
		return nil
	}
	if err := os.Remove(path); err != nil {
		// A non-empty directory is expected when it contains skipped files.
		if errors.Is(err, syscall.ENOTEMPTY) || errors.Is(err, syscall.EEXIST) {
			return nil
		}
		return err
	}
	return nil
}
