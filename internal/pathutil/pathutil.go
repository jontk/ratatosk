// Package pathutil provides small filesystem path helpers shared across
// ratatosk packages.
package pathutil

import (
	"os"
	"strings"
)

// ExpandTilde replaces a leading "~" or "~/" with the user's home directory
// (from $HOME). Bare "~user" forms and paths without a leading tilde are
// returned unchanged. If $HOME is unset, the path is returned as-is.
func ExpandTilde(path string) string {
	if path == "~" || strings.HasPrefix(path, "~/") {
		if home := os.Getenv("HOME"); home != "" {
			return home + path[1:]
		}
	}
	return path
}
