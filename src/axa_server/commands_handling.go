package axa_server

import (
  "fmt"
  "log"
)

import (
  "axaDB/src/dberrs"
  "axaDB/src/axa_security"
  "axaDB/src/fs"
  "axaDB/src/parsers"
)

func fetch(split []string) string{

	collection,url,_,err := fetchExecParams("from", split)
	if err != "" {
		return err
	}

	if len(url) == 0 {
		return handleNoUrlFetch(collection)
	} 
	return handleUrlFetch(collection, url)
}

func feed(split []string) (ret string){
	defer func() {
		if err := recover(); err != nil {
			ret = dberrs.DB_EX06().Err
		}
		ret = dberrs.DB_NORM().Err
	}()

	collection, url, json, err := fetchExecParams("in", split)
	if err != "" {
		return err
	}

	jsonMap, newErr := parsers.JsonMapFromString(json)
	if newErr != nil {
		return dberrs.DB_EX05().Err
	}

	dataFile := fs.FindDataFileContainingKey(collection, string(url[0][0]))
	if dataFile == "" {
		fmt.Println("(executioner) new file")
		dataFile = "./" + collection + "/" + string(url[0][0]) + ".db"

		// TODO: handle error
		file, _ := fs.CreateFile(dataFile)

		jsonMarsh := parsers.CreateFeedMarhsall(url, jsonMap)
 
		ret := fs.WriteToEmptyFile(file, jsonMarsh).Err
		file.Close()
		return ret
	} else {
		fmt.Println("(executioner) not new file")
		oldMap := fs.JsonMapFromFile(dataFile)
		newMap := make(map[string]interface{})
		iterMap := oldMap
		//for _, el := range url {
		for key, value := range iterMap {
			if key != url[0] {
				newMap[key] = value
			}
		}
		iterMap = iterMap[url[0]].(map[string]interface{})
		newMap[url[0]] = diveNewMap(url[1:], iterMap, jsonMap)

		fs.WriteMapToFile(newMap, dataFile)
		return fmt.Sprintf("%s", newMap)	
	}

	log.Println(dataFile)
	
	return "(axa executioner): feed succesful!"
}

func delete(split []string) string{
	_,_,_,err := fetchExecParams("from", split)
	if err != "" {
		return err
	}
	return "(axa executioner): deletion succesful!"
}

func login(split []string) string{
  login := split[0]
  loginkey := string(login[0])
  password := split[1]
  dataFile := fs.FindDataFileContainingKey("AXA_USERS", loginkey)
  if dataFile == "" {
    return ""
    // TODO : create axa err for this (data file with key in name not found)
  }
  usersMap := fs.JsonMapFromFile(dataFile)
  if axa_security.SHA256EncryptPassword(password) == usersMap[login].(map[string]interface{})["password"] {
    return "(axa execution): login success"
  } else {
    return "(axa execution): login failed"
  }
}


