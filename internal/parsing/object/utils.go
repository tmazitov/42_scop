package objectParsing 

import (
	"strings"
	"path/filepath"
)

func isPath(s string) bool {
    // Check for common path separators
    return strings.Contains(s, "/") || 
           strings.Contains(s, "\\") ||
           strings.Contains(s, string(filepath.Separator))
}