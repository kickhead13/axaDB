package axa_server

import(
	"fmt"
)

import (
	"axaDB/src/parsers"
	"axaDB/src/dberrs"
	"axaDB/src/fs"
)

func fetchExecParams(colKeyword string, split []string) (string, []string, string, string){
	fmt.Println("test")
	collection, err := parsers.CollectionFromSplit(split, colKeyword)
	if err != dberrs.DB_NORM() {
		return "",[]string{},"",err.Err
	}

	url, err := parsers.UrlFromSplit(split)
	if err != dberrs.DB_NORM() {
		return "",[]string{},"",err.Err
	}

	json, err := parsers.JsonFromSplit(split)
	if err != dberrs.DB_NORM() {
		return "",[]string{},"",err.Err
	}

	return collection, url, json, ""
}

func diveNewMap(url []string, iterMap map[string]interface{}, jsonMap map[string]interface{}) map[string]interface{} {
	if len(url) == 0 {
		return jsonMap
	}
	if iterMap == nil {
		return jsonMap
	}
	newMap := make(map[string]interface{})
	for key, value := range iterMap {
		if key != url[0] {
			newMap[key] = value
		}
	}
	iterMap = iterMap[url[0]].(map[string]interface{})
	newMap[url[0]] = diveNewMap(url[1:], iterMap, jsonMap)
	return newMap
}

func handleNoUrlFetch(collection string) string {
	return ""
}

func diveForFetch(dataFile string, url []string) (str string) {
	defer func() {
		if err := recover(); err != nil {
			str = ""
		}
	}()
	dfMap := fs.JsonMapFromFile(dataFile)
	retMap := dfMap
	for _, el := range url {
		retMap = retMap[el].(map[string]interface{})
	}
	return fmt.Sprintf("%s", retMap)
}

func handleUrlFetch(collection string, url []string) string{
	if ! fs.CollectionExists(collection) {
		return ""
	}
	dataFile := fs.FindDataFileContainingKey(collection, string(url[0][0]))
	if dataFile == "" {
		return ""
	}
	return diveForFetch(dataFile, url)
}