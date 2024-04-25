package fs

import (
	"os"
	"fmt"
	"axaDB/src/axa_security"
)

import (
	"axaDB/src/dberrs"
)

func CreateInitFile(dir string, cpuCores string, possibleBackups string, maxDataFileSize string, databaseName string) dberrs.AxaErr{
	fmt.Println("(axa init) creating init.cfg for database at " + dir + "...")

	f,err := os.Create(dir+"init.cfg")
	defer f.Close()
	if err != nil {
		return dberrs.DB_D04()
	}

	fileContent := `{
  		"cpuCores":"` + cpuCores +`",
   		"possibleBackups":"` + possibleBackups +`",
   		"maxDataFileSize":"` + maxDataFileSize + `",
   		"databaseName":"` + databaseName + `"
	}`
	encryptedFileContent, err := axa_security.EncryptData(fileContent)
	if err != nil {
		return dberrs.DB_E01()
	}

	f.WriteString(encryptedFileContent)
	return dberrs.DB_NORM()
}

func CreateUsersDefaultDataFile(dir string, sys_password string) dberrs.AxaErr{
	fmt.Println("(axa init) creating AXA_USERS collection default data file...")

	datafilePath := dir+"AXA_USERS/df.db"
	f,err := os.Create(datafilePath)
	defer f.Close()
	if err != nil {
		return dberrs.DB_D05()
	}

	fileContent := `{
		"sys":{
			"password":"`+axa_security.SHA256EncryptPassword(sys_password)+`",
			"role":"AXA_ADMIN"
		}
	}`
	encryptedFileContent, err := axa_security.EncryptData(fileContent)
	if err != nil {
		return dberrs.DB_E01()
	}
	f.WriteString(encryptedFileContent)

	return dberrs.DB_NORM()
}

func CreateCollectionRulesFile(dir string, collection string, rules map[string]string) dberrs.AxaErr{

	fileContent := "{"
	iter := 0
	for role, perms := range rules {
		if iter == 0 {
			fileContent += "\n   \"" + role + "\":\"" + perms + "\""
		} else {
			fileContent += ",\n   \"" + role + "\":\"" + perms + "\""
		}
		iter += 1
	}
	fileContent += "\n}"

	datafilePath := dir+collection+"/rules.db"
	f,err := os.Create(datafilePath)
	defer f.Close()
	if err != nil {
		return dberrs.DB_D06(collection)
	}

	encryptedFileContent, err := axa_security.EncryptData(fileContent)
	if err != nil {
		return dberrs.DB_E01()
	}
	f.WriteString(encryptedFileContent)

	return dberrs.DB_NORM()
}