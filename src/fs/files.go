package fs

import (
	"os"
)

func CreateInitFile(dir string, cpuCores string, possibleBackups string, maxDataFileSize string) {
	f,_ := os.Create(dir+"init.cfg")
	defer f.Close()
	f.WriteString("{")
	f.WriteString("\n\t\"cpuCores\":" + cpuCores +",")
	f.WriteString("\n\t\"possibleBackups\":" + possibleBackups + ",")
	f.WriteString("\n\t\"maxDataFileSize\":" + maxDataFileSize + "")
	f.WriteString("\n}")
}