package app

import (
	"path/filepath"
	"strings"
)

// EnsureRoot will append 'prefix' to 'path' if necessary.
func EnsureRoot(path string, prefix string) string {

	path = strings.TrimLeft(path, "/")

	if prefix == "" {
		return "/" + path
	}

	path = filepath.Join(prefix, path)
	return path
}
