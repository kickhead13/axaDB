package parsers

import (
	"encoding/json"
)

func JsonMapFromString(str string) (map[string]interface{}, error) {
	var payload map[string]interface{}
	err := json.Unmarshal([]byte(str), &payload)
	return payload, err
}

func StringFromJsonMap(mMap map[string]interface{}) string {
	bContent, _ := json.Marshal(mMap)
	return string(bContent)
}

func nextHead(url []string, js map[string]interface{}) map[string]interface{}{
	if len(url) <= 0 {
		return js
	}
	head := make(map[string]interface{})
	head[url[0]] = nextHead(url[1:], js)
	return head
}

func CreateFeedMarhsall(url []string, js map[string]interface{}) string{
	head := make(map[string]interface{})
	head[url[0]] = nextHead(url[1:], js)
	content, _ := json.Marshal(head)
	return string(content)
}