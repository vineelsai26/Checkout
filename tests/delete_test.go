package tests

import (
	"testing"

	"vineelsai.com/checkout/src/utils"
)

func TestDelete(t *testing.T) {
	err := utils.DeleteFolder("testdir")
	if err != nil {
		t.Errorf("error deleting directory: %s", err)
	}
}
