package ya2

import (
	"strings"
	"testing"
)

func TestPrintAllFiles(t *testing.T) {
	PrintAllFiles("/Users/a.f.kay/Code/ya-go")
}

func TestPrintAllFilesWithFilter(t *testing.T) {
	PrintAllFilesWithFilter("/Users/a.f.kay/Code/ya-go", "go.mod")
}

func TestPrintFilesRecursive(t *testing.T) {
	PrintFilesRecursive(".", func(s string) bool { return strings.Contains(s, "go.mod") })
}
