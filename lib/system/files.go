package system

import (
	"os"
	"strings"
)

func FilenameFromPath(path string) string {
	return path[strings.LastIndex(path, "/")+1:]
}

func LocalExists(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	} else {
		return false
	}
}
