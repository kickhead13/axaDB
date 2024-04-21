package axa_init

import (
	"fmt"
	"axaDB/src/parsers"
	"axaDB/src/fs"
	"os"
	"axaDB/src/dberrs"
	"strings"
)

func Init(args []string) (dberrs.AxaErr) {
	if len(args) == 0 {
		fmt.Println(parsers.InitHelp())
	} else {
		
		at, _ := parsers.InitParse(args, []string{"--at", "-@"}[:])
		at = fs.FormatDirName(at)
		if _, err := os.Stat(at); !os.IsNotExist(err) {
			return dberrs.DB_D01()
		} else {
			err = os.Mkdir(at, 0755)
			if err != nil {
				dberrs.DB_D02()
			}
		}

		fmt.Println("(axa init) created database directory successfuly...")

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

		databaseName, err := parsers.InitParse(args, []string{"--databaseName", "-dbn"}[:])
		if err != nil {
			databaseName = strings.Split(at, "/")[len(strings.Split(at, "/"))-2]
		}
		aerr := fs.CreateInitFile(at, cpuCores, possibleBackups, maxDataFileSize, databaseName)
		if aerr != dberrs.DB_NORM() {
			return aerr
		}

		sysPassword, err := parsers.InitParse(args, []string{"--sysPassword", "-sp"}[:])
		if err != nil {
			sysPassword = "veryBadPassword"
		}

		fmt.Println("(axa init) creating AXA_USERS collection...")
		err = os.Mkdir(at + "AXA_USERS", 0755)
		if err != nil {
			return dberrs.DB_D03()
		}

		fmt.Println("(axa init) creating AXA_USERS default data file containing sys user info...")
		aerr = fs.CreateUsersDefaultDataFile(at, sysPassword)
		if aerr != dberrs.DB_NORM() {
			return aerr
		}

		aerr = fs.CreateCollectionRulesFile(at, "AXA_USERS", map[string]string{
			"AXA_ADMIN":"read|write|modify",
			"#":"|",
		})
	}
	return dberrs.DB_NORM()
}

