package fs

import (
	"strings"
	"os"
)

func FormatDirName(dir string) string{
	if strings.HasSuffix(dir, "/") {
		return dir
	} 
	return dir + "/"
}

func CollectionExists(collection string) bool{
	_, err := os.Stat("./" + collection)
	if err == nil {
		return true
	}
	return false
}