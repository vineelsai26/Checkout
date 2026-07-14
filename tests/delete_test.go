package tests

import (
	"os"
	"path/filepath"
	"testing"

	"vineelsai.com/checkout/src/utils"
)

func TestDelete(t *testing.T) {
	root := t.TempDir()
	targetDir := filepath.Join(root, "target")

	if err := os.MkdirAll(targetDir, 0755); err != nil {
		t.Fatalf("create target directory: %s", err)
	}

	if err := utils.DeleteFolder(targetDir); err != nil {
		t.Fatalf("delete directory: %s", err)
	}

	if utils.Exists(targetDir) {
		t.Fatalf("target directory still exists after delete")
	}
}

func TestDeleteCopiedTreePreservesSkippedFiles(t *testing.T) {
	root := t.TempDir()
	targetDir := filepath.Join(root, "target")
	if err := os.MkdirAll(filepath.Join(targetDir, "src"), 0755); err != nil {
		t.Fatalf("create target directory: %s", err)
	}
	if err := os.WriteFile(filepath.Join(targetDir, "src", "main.go"), []byte("package main"), 0644); err != nil {
		t.Fatalf("write copied file: %s", err)
	}
	if err := os.WriteFile(filepath.Join(targetDir, ".env"), []byte("SECRET=test"), 0600); err != nil {
		t.Fatalf("write skipped file: %s", err)
	}

	if err := utils.DeleteCopiedTree(targetDir, utils.CopyOptions{}); err != nil {
		t.Fatalf("delete copied tree: %s", err)
	}

	if utils.Exists(filepath.Join(targetDir, "src", "main.go")) {
		t.Fatalf("copied file still exists")
	}
	if !utils.Exists(filepath.Join(targetDir, ".env")) {
		t.Fatalf("skipped .env file was deleted")
	}
}
