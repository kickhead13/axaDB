package fs

import (
	"os"
)

import (
	"axaDB/src/axa_security"
	"axaDB/src/parsers"
)

func JsonMapFromFile(dataFile string) map[string]interface{}{

	// TODO: handle error
	data, _ := os.ReadFile(dataFile)
	decryptedBytes, _ := axa_security.DecryptData(string(data))
	retMap, _ := parsers.JsonMapFromString(string(decryptedBytes))
	return retMap

}