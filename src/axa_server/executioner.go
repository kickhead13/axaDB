package axa_server

import (
	"time"
	"math/rand/v2"
	"sync"
	"strings"
	"log"
	"fmt"
)

import (
	"axaDB/src/fs"
	"axaDB/src/dberrs"
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
func feed(split []string) string{
	defer func() string{
		if err := recover(); err != nil {
			return dberrs.DB_EX06().Err
		}
		return ""
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

func execute(cmds map[string]string, responseBuffer *CritBuffer) {
	var response string
	for user, cmd := range cmds {
		split := strings.Split(cmd, " ")
		switch split[0] {
		case "feed":
			response = feed(split[1:])
		case "fetch":
			response = fetch(split[1:])
		case "delete":
			response = delete(split[1:])
		default:
			response = dberrs.DB_EX01().Err
		}
		responseBuffer.push(response, user)
	}
	
}

func handleExecutioner(execBuffer *CritBuffer, responseBuffer *CritBuffer, wgroup *sync.WaitGroup) {

	for {
		time.Sleep(time.Duration(rand.IntN(100)) * time.Millisecond)
		//execBuffer.print()
		commands := execBuffer.read()
		if len(commands) > 0 {
			execute(commands, responseBuffer)
		} 
	}
	wgroup.Done()
}