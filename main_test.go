package main

import (
	"os"
	"path/filepath"
	"testing"

	"vineelsai.com/checkout/src/utils"
)

func TestRunInitDryRunUsesConfiguredRoots(t *testing.T) {
	originalSourceDir := utils.ProjectSourceDir
	originalCheckoutRoot := utils.ProjectCheckoutRootDir
	t.Cleanup(func() {
		utils.ProjectSourceDir = originalSourceDir
		utils.ProjectCheckoutRootDir = originalCheckoutRoot
	})

	root := t.TempDir()
	sourceRoot := filepath.Join(root, "source")
	checkoutRoot := filepath.Join(root, "checkout")
	projectDir := filepath.Join(sourceRoot, "vineelsai26", "sample-project")

	if err := os.MkdirAll(filepath.Join(projectDir, "src"), 0755); err != nil {
		t.Fatalf("create project: %s", err)
	}
	if err := os.WriteFile(filepath.Join(projectDir, "src", "index.js"), []byte("test"), 0644); err != nil {
		t.Fatalf("write source file: %s", err)
	}
	if err := os.MkdirAll(checkoutRoot, 0755); err != nil {
		t.Fatalf("create checkout root: %s", err)
	}

	err := run([]string{
		"init",
		"--dry-run",
		"--no-open",
		"--source-dir", sourceRoot,
		"--checkout-root", checkoutRoot,
		"sample-project",
	})
	if err != nil {
		t.Fatalf("run init dry-run: %s", err)
	}

	if !utils.Exists(projectDir) {
		t.Fatalf("dry-run deleted source project")
	}
	if utils.Exists(filepath.Join(checkoutRoot, "sample-project")) {
		t.Fatalf("dry-run created checkout project")
	}
}

func TestRunDeinitRejectsTraversal(t *testing.T) {
	originalSourceDir := utils.ProjectSourceDir
	originalCheckoutRoot := utils.ProjectCheckoutRootDir
	t.Cleanup(func() {
		utils.ProjectSourceDir = originalSourceDir
		utils.ProjectCheckoutRootDir = originalCheckoutRoot
	})

	root := t.TempDir()
	utils.ProjectSourceDir = filepath.Join(root, "source")
	utils.ProjectCheckoutRootDir = filepath.Join(root, "checkout")

	if err := run([]string{"deinit", "--source-folder", "../../outside", "sample-project"}); err == nil {
		t.Fatal("expected traversal error")
	}
}
