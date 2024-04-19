package axa_init

import (
	"fmt"
	"axaDB/src/parsers"
	"axaDB/src/fs"
	"os"
	"errors"
)

func Init(args []string) (error) {
	if len(args) == 0 {
		fmt.Println(parsers.InitHelp())
	} else {
		
		at, _ := parsers.InitParse(args, []string{"--at", "-a"}[:])
		at = fs.FormatDirName(at)
		if _, err := os.Stat(at); !os.IsNotExist(err) {
			return errors.New("(axa init) database directory: alreadry exists \n\t& database directory must be created at database creation time")
		} else {
			_ = os.Mkdir(at, 0755)
		}


		cpuCores, err := parsers.InitParse(args, []string{"--cpuCores", "-cc"}[:])
		if err != nil {
			cpuCores = "4"
		}

		possibleBackups, err := parsers.InitParse(args, []string{"--possibleBackups", "-pb"}[:])
		if err != nil {
			possibleBackups = "4"
		}

		maxDataFileSize, err := parsers.InitParse(args, []string{"--maxDataFileSize", "-mdf"}[:])
		if err != nil {
			maxDataFileSize = "1024"
		}
		fs.CreateInitFile(at, cpuCores, possibleBackups, maxDataFileSize)

	}
	return nil
}

