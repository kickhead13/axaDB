package axa_server

import (
	"axaDB/src/parsers"
	"axaDB/src/dberrs"
)

func fetchExecParams(colKeyword string, split []string) (string, []string, string, string){
	collection, err := parsers.CollectionFromSplit(split, "in")
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