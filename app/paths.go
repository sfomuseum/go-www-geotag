package app

import (
	"fmt"
	"path/filepath"
	"strings"
)

// EnsureRoot will append 'prefix' to 'path' if necessary.
func EnsureRoot(path string, prefix string) string {

	path = strings.TrimLeft(path, "/")

	if prefix == "" {
		return "/" + path
	}

	trailing_slash := strings.HasSuffix(path, "/")

	path = filepath.Join(prefix, path)

	if trailing_slash && !strings.HasSuffix(path, "/") {
		path = fmt.Sprintf("%s/", path)
	}

	return path
}
