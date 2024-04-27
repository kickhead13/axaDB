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
		messages: map[string]string{},
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

func (buffer *CritBuffer) print() {
	buffer.locket.Lock()
	defer buffer.locket.Unlock()
	for user, cmd := range buffer.messages {
		fmt.Printf("%s : %s\n", user, cmd)
	}
}
