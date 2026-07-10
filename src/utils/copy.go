package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type CopyOptions struct {
	IncludeEnv bool
	DryRun     bool
	Excludes   []string
}

var DefaultExcludedNames = []string{
	"node_modules",
	"bower_components",
	".next",
	".nuxt",
	".svelte-kit",
	"dist",
	"build",
	"out",
	"coverage",
	".turbo",
	".cache",
	".parcel-cache",
	".vite",
	".venv",
	"venv",
	"env",
	"__pycache__",
	".pytest_cache",
	".mypy_cache",
	".ruff_cache",
	".tox",
	".nox",
	"target",
}

func (options CopyOptions) ShouldSkip(path string) bool {
	name := filepath.Base(path)
	if isEnvFile(name) && !options.IncludeEnv {
		return true
	}

	for _, exclude := range append(DefaultExcludedNames, options.Excludes...) {
		exclude = strings.TrimSpace(exclude)
		if exclude == "" {
			continue
		}
		if name == exclude || pathMatchesPattern(path, exclude) {
			return true
		}
	}

	return false
}

func isEnvFile(name string) bool {
	if name == ".env.example" || name == ".env.template" || name == ".env.sample" {
		return false
	}
	return name == ".env" || strings.HasPrefix(name, ".env.")
}

func pathMatchesPattern(path string, pattern string) bool {
	if !strings.ContainsAny(pattern, "*?[") {
		return false
	}

	name := filepath.Base(path)
	if matched, err := filepath.Match(pattern, name); err == nil && matched {
		return true
	}
	if matched, err := filepath.Match(pattern, path); err == nil && matched {
		return true
	}

	return false
}

func CopyDirectory(sourceDir, dest string) error {
	return CopyDirectoryWithOptions(sourceDir, dest, CopyOptions{})
}

func CopyDirectoryWithOptions(sourceDir, dest string, options CopyOptions) error {
	if options.ShouldSkip(sourceDir) {
		fmt.Println("Skipping", sourceDir)
		return nil
	}

	fmt.Println("Copying", sourceDir, "to", dest)
	entries, err := os.ReadDir(sourceDir)
	if err != nil {
		return err
	}

	if !options.DryRun {
		if err := CreateIfNotExists(dest, 0755); err != nil {
			return err
		}
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(sourceDir, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		if options.ShouldSkip(sourcePath) {
			fmt.Println("Skipping", sourcePath)
			continue
		}

		if err := copyEntry(sourcePath, destPath, options); err != nil {
			return err
		}
	}
	return nil
}

func copyEntry(sourcePath, destPath string, options CopyOptions) error {
	fileInfo, err := os.Lstat(sourcePath)
	if err != nil {
		return err
	}

	if options.DryRun {
		fmt.Println("Would copy", sourcePath, "to", destPath)
		return nil
	}

	switch fileInfo.Mode() & os.ModeType {
	case os.ModeDir:
		if err := CreateIfNotExists(destPath, 0755); err != nil {
			return err
		}
		if err := CopyDirectoryWithOptions(sourcePath, destPath, options); err != nil {
			return err
		}
	case os.ModeSymlink:
		if err := CopySymLink(sourcePath, destPath); err != nil {
			return err
		}
	default:
		if err := Copy(sourcePath, destPath); err != nil {
			return err
		}
	}

	if err := preserveOwnership(sourcePath, destPath, fileInfo); err != nil {
		return err
	}

	isSymlink := fileInfo.Mode()&os.ModeSymlink != 0
	if !isSymlink {
		return os.Chmod(destPath, fileInfo.Mode())
	}

	return nil
}

func Copy(srcFile, dstFile string) error {
	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}

	defer in.Close()

	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}

func CopySymLink(source, dest string) error {
	link, err := os.Readlink(source)
	if err != nil {
		return err
	}
	return os.Symlink(link, dest)
}
