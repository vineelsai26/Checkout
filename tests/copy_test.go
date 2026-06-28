package tests

import (
	"os"
	"path/filepath"
	"testing"

	"vineelsai.com/checkout/src/utils"
)

func TestCopyDir(t *testing.T) {
	root := t.TempDir()
	sourceDir := filepath.Join(root, "source")
	destDir := filepath.Join(root, "dest")

	if err := os.MkdirAll(filepath.Join(sourceDir, "nested"), 0755); err != nil {
		t.Fatalf("create source directory: %s", err)
	}
	if err := os.WriteFile(filepath.Join(sourceDir, "nested", "test.txt"), []byte("checkout"), 0644); err != nil {
		t.Fatalf("create source file: %s", err)
	}

	if err := utils.CopyDirectory(sourceDir, destDir); err != nil {
		t.Fatalf("copy directory: %s", err)
	}

	copied, err := os.ReadFile(filepath.Join(destDir, "nested", "test.txt"))
	if err != nil {
		t.Fatalf("read copied file: %s", err)
	}
	if string(copied) != "checkout" {
		t.Fatalf("copied file content = %q, want %q", copied, "checkout")
	}
}

func TestCopy(t *testing.T) {
	root := t.TempDir()
	sourceFile := filepath.Join(root, "source.txt")
	destFile := filepath.Join(root, "dest.txt")

	if err := os.WriteFile(sourceFile, []byte("checkout"), 0644); err != nil {
		t.Fatalf("create source file: %s", err)
	}

	if err := utils.Copy(sourceFile, destFile); err != nil {
		t.Fatalf("copy file: %s", err)
	}

	copied, err := os.ReadFile(destFile)
	if err != nil {
		t.Fatalf("read copied file: %s", err)
	}
	if string(copied) != "checkout" {
		t.Fatalf("copied file content = %q, want %q", copied, "checkout")
	}
}
