package utils

import "testing"

func TestDelete(t *testing.T) {
	err := DeleteFolder("testdir")
	if err != nil {
		t.Errorf("error deleting directory: %s", err)
	}
}
