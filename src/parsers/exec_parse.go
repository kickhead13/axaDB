package parsers

import "axaDB/src/dberrs"

func CollectionFromSplit(split []string, statement string) (string, dberrs.AxaErr){

	var inCount bool = false
	for _, element := range split {
		if element == statement {
			if inCount == true {
				return "", dberrs.DB_EX02(statement)
			}
			inCount = true
		} 
	}

	for index, element := range split {
		if index > 0 {
			if split[index-1] == statement {
				return element, dberrs.DB_NORM()
			}
		}
	}
	
	return "", dberrs.DB_EX03()
}

func UrlFromSplit(split []string) ([]string, dberrs.AxaErr){
	url := []string{}
	for index, element := range split {
		if index > 0 {
			if split[index-1] == "at" && element == "at" {
				return []string{}, dberrs.DB_EX04()
			}
			if split[index-1] == "at" {
				url = append(url, element)
			}
		}
	}
	return url, dberrs.DB_NORM()
}

func findLastAt(split []string) int{
	var pos int = 0
	for index, _ := range split {
		if index > 0 {
			if split[index - 1] == "at" {
				pos = index
			}
		}
	}
	return pos
}

func JsonFromSplit(split []string) (string, dberrs.AxaErr) {
	var json string
	pos := findLastAt(split)
	for _, element := range split[(pos+1):] {
		json += element
	}
	return json, dberrs.DB_NORM()
}
