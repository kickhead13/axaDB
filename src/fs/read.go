package fs

import (
	"os"
  "fmt"
)

import (
	"axaDB/src/axa_security"
	"axaDB/src/parsers"
)

func JsonMapFromFile(dataFile string) map[string]interface{}{

	// TODO: handle error
	data, err := os.ReadFile(dataFile)
  if err != nil {
    fmt.Println(err)
  }
	decryptedBytes, err := axa_security.DecryptData(string(data))
  if err != nil {
    fmt.Println(err)
  }
	retMap, err := parsers.JsonMapFromString(string(decryptedBytes))
  if err != nil {
    fmt.Println(err)
  }
	return retMap

}
