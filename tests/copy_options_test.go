package tests

import (
	"os"
	"path/filepath"
	"testing"

	"vineelsai.com/checkout/src/utils"
)

func TestCopyDirectoryExcludesLocalDependenciesAndCaches(t *testing.T) {
	root := t.TempDir()
	sourceDir := filepath.Join(root, "source")
	destDir := filepath.Join(root, "dest")

	writeFile(t, filepath.Join(sourceDir, "src", "main.ts"), "app")
	writeFile(t, filepath.Join(sourceDir, "node_modules", "dep", "index.js"), "dependency")
	writeFile(t, filepath.Join(sourceDir, "__pycache__", "module.pyc"), "cache")
	writeFile(t, filepath.Join(sourceDir, ".venv", "bin", "python"), "python")

	if err := utils.CopyDirectoryWithOptions(sourceDir, destDir, utils.CopyOptions{}); err != nil {
		t.Fatalf("copy directory: %s", err)
	}

	assertExists(t, filepath.Join(destDir, "src", "main.ts"))
	assertNotExists(t, filepath.Join(destDir, "node_modules"))
	assertNotExists(t, filepath.Join(destDir, "__pycache__"))
	assertNotExists(t, filepath.Join(destDir, ".venv"))
}

func TestCopyDirectoryExcludesEnvFilesByDefault(t *testing.T) {
	root := t.TempDir()
	sourceDir := filepath.Join(root, "source")
	destDir := filepath.Join(root, "dest")

	writeFile(t, filepath.Join(sourceDir, ".env"), "SECRET=true")
	writeFile(t, filepath.Join(sourceDir, ".env.local"), "SECRET=true")
	writeFile(t, filepath.Join(sourceDir, ".env.example"), "SECRET=false")

	if err := utils.CopyDirectoryWithOptions(sourceDir, destDir, utils.CopyOptions{}); err != nil {
		t.Fatalf("copy directory: %s", err)
	}

	assertNotExists(t, filepath.Join(destDir, ".env"))
	assertNotExists(t, filepath.Join(destDir, ".env.local"))
	assertExists(t, filepath.Join(destDir, ".env.example"))
}

func TestCopyDirectoryCanIncludeEnvFiles(t *testing.T) {
	root := t.TempDir()
	sourceDir := filepath.Join(root, "source")
	destDir := filepath.Join(root, "dest")

	writeFile(t, filepath.Join(sourceDir, ".env"), "SECRET=true")
	writeFile(t, filepath.Join(sourceDir, ".env.local"), "SECRET=true")

	if err := utils.CopyDirectoryWithOptions(sourceDir, destDir, utils.CopyOptions{IncludeEnv: true}); err != nil {
		t.Fatalf("copy directory: %s", err)
	}

	assertExists(t, filepath.Join(destDir, ".env"))
	assertExists(t, filepath.Join(destDir, ".env.local"))
}

func TestCopyDirectorySupportsAdditionalExcludes(t *testing.T) {
	root := t.TempDir()
	sourceDir := filepath.Join(root, "source")
	destDir := filepath.Join(root, "dest")

	writeFile(t, filepath.Join(sourceDir, "keep.txt"), "keep")
	writeFile(t, filepath.Join(sourceDir, "tmp", "drop.txt"), "drop")
	writeFile(t, filepath.Join(sourceDir, "debug.log"), "drop")

	options := utils.CopyOptions{Excludes: []string{"tmp", "*.log"}}
	if err := utils.CopyDirectoryWithOptions(sourceDir, destDir, options); err != nil {
		t.Fatalf("copy directory: %s", err)
	}

	assertExists(t, filepath.Join(destDir, "keep.txt"))
	assertNotExists(t, filepath.Join(destDir, "tmp"))
	assertNotExists(t, filepath.Join(destDir, "debug.log"))
}

func writeFile(t *testing.T, path string, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatalf("create parent directory: %s", err)
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("write test file: %s", err)
	}
}

func assertExists(t *testing.T, path string) {
	t.Helper()
	if !utils.Exists(path) {
		t.Fatalf("expected %s to exist", path)
	}
}

func assertNotExists(t *testing.T, path string) {
	t.Helper()
	if utils.Exists(path) {
		t.Fatalf("expected %s not to exist", path)
	}
}
