package axa_server

import (
	"time"
	"math/rand/v2"
	"sync"
	"strings"
)

import (
	"axaDB/src/dberrs"
)

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
    case "login":
      response = login(split[1:])
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
