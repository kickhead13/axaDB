package axa_server

import (
	"time"
	"math/rand/v2"
	"sync"
)

func execute(cmds map[string]string) {
	for _, cmd := range cmds {
		cmd = cmd
	}
	
}

func handleExecutioner(execBuffer *CritBuffer, responseBuffer *CritBuffer, wgroup *sync.WaitGroup) {

	for {
		time.Sleep(time.Duration(rand.IntN(100)) * time.Millisecond)
		//execBuffer.print()
		commands := execBuffer.read()
		if len(commands) > 0 {
			execute(commands)
		} 
	}
	wgroup.Done()
}