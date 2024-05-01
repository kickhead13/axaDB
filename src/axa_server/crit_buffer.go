package axa_server

import (
	"sync"
	"fmt"
)

type CritBuffer struct {
	locket sync.Mutex
	messages map[string]string
}

func InitCritBuffer() CritBuffer{
	return CritBuffer{
		messages: make(map[string]string),
	}
}

func (buffer *CritBuffer) push(exec string, user string) {
	buffer.locket.Lock()
	defer buffer.locket.Unlock()
	buffer.messages[user] = exec
}

func (buffer *CritBuffer) read() map[string]string {
	buffer.locket.Lock()
	defer buffer.locket.Unlock()
	messages := buffer.messages
	buffer.messages = map[string]string{}
	return messages
}

func mDelete(pMap *map[string]string, rKey string) {
	nMap := make(map[string]string)
	for key, value := range *pMap {
		if key != rKey {
			nMap[key] = value
		}
	}
	*pMap = nMap
}

func (buffer *CritBuffer) readValueOfKey(key string) string{
	buffer.locket.Lock()
	defer buffer.locket.Unlock()
	if val, ok := buffer.messages[key]; ok {
		mDelete(&buffer.messages, key)
		return val
	}
	return ""
}

func (buffer *CritBuffer) print() {
	buffer.locket.Lock()
	defer buffer.locket.Unlock()
	for user, cmd := range buffer.messages {
		fmt.Printf("%s : %s\n", user, cmd)
	}
}
