package tests

import (
	"os/exec"
	"testing"

	"vineelsai.com/checkout/src/utils"
)

func TestCopyDir(t *testing.T) {
	err := utils.CopyDirectory("testdir", "testdir1")
	if err != nil {
		t.Errorf("error copying directory: %s", err)
	}

	command := "diff --brief -r testdir testdir1"

	out, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		t.Errorf("error folders are not the same: %s", err)
	}

	if len(out) > 0 {
		t.Errorf("error folders are not the same: %s", out)
	}
}

func TestCopy(t *testing.T) {
	err := utils.Copy("testfile", "testfile1")
	if err != nil {
		t.Errorf("error copying directory: %s", err)
	}
}
