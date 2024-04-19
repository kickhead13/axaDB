package fs

import (
	"os"
	"fmt"
)

func CreateInitFile(dir string, cpuCores string, possibleBackups string, maxDataFileSize string) {
	fmt.Println("(axa init) creating init.cfg for database at " + dir + "...")

	f,_ := os.Create(dir+"init.cfg")
	defer f.Close()
	f.WriteString("{")
	f.WriteString("\n\t\"cpuCores\":" + cpuCores +",")
	f.WriteString("\n\t\"possibleBackups\":" + possibleBackups + ",")
	f.WriteString("\n\t\"maxDataFileSize\":" + maxDataFileSize + "")
	f.WriteString("\n}")
}