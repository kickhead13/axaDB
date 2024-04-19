package fs

import (
	"strings"
)

func FormatDirName(dir string) string{
	if strings.HasSuffix(dir, "/") {
		return dir
	} 
	return dir + "/"
}