package testhelpers

import (
	"os"
	"testing"
)

func setupTest(t *testing.T) {
	os.Chdir(os.Getenv("WORKSPACE_DIR"))
}

func Setup(t *testing.T) {
	setupTest(t)
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
