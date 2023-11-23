package testhelpers

import (
	"os"
	"testing"
)

func Setup(t *testing.T) {
	err := os.Chdir(os.Getenv("WORKSPACE_DIR"))
	if err != nil {
		return
	}
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

func LoadJSON(path string) string {
	jsonFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(jsonFile)
}
