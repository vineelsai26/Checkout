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
