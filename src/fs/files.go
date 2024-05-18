package fs

import (
	"os"
	"fmt"
	"strings"
)

import (
	"axaDB/src/dberrs"
	"axaDB/src/axa_security"
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

	datafilePath := dir+"AXA_USERS/s.db"
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

	datafilePath := dir+collection+"/rules.axa"
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

func FindDataFileContainingKey(collection string, key string) string{
	entries, err := os.ReadDir("./" + collection)
  if err != nil {
	fmt.Println(err)
    return ""
  } 
  for _, e := range entries {
    if strings.Contains(strings.Replace(e.Name(), ".db", "", -1), string(key[0])) && strings.Contains(e.Name(), ".db"){
      return "./" + collection + "/" + e.Name()
    }
  }
	
	return ""
}

func CreateFile(path string) (*os.File, error){
	f, err := os.Create(path)
	return f, err
}

func WriteToEmptyFile(file *os.File, str string) dberrs.AxaErr{
	encryptedStr, err := axa_security.EncryptData(str)
	if err != nil {
		return dberrs.DB_E01()
	}
	file.WriteString(encryptedStr)
	fmt.Println("(axa executioner) fed collection")
	return dberrs.DB_FED()
}

func WriteMapToFile(mMap map[string]interface{}, file string) {
	_ = os.Remove(file)
	f, _ := os.Create(file)
	_ = WriteToEmptyFile(f, fmt.Sprintf("%s", mMap))
	f.Close()
}

func WriteToUnemptyFile(dataFile string, str string, url []string) dberrs.AxaErr{
	return dberrs.DB_NORM()	
}
