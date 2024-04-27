package axa_server

import (
	"time"
	"math/rand/v2"
	"sync"
	"strings"
	"axaDB/src/dberrs"
	"axaDB/src/parsers"
)

func fetch(split []string) string{
	_, err := parsers.CollectionFromSplit(split, "from")
	if err != dberrs.DB_NORM() {
		return err.Err
	}

	_, err = parsers.UrlFromSplit(split)
	if err != dberrs.DB_NORM() {
		return err.Err
	}

	return ""
}

func feed(split []string) string{
	_, err := parsers.CollectionFromSplit(split, "in")
	if err != dberrs.DB_NORM() {
		return err.Err
	}

	_, err = parsers.UrlFromSplit(split)
	if err != dberrs.DB_NORM() {
		return err.Err
	}

	_, err = parsers.JsonFromSplit(split)
	if err != dberrs.DB_NORM() {
		return err.Err
	}

	return ""
}

func delete(split []string) string{
	_, err := parsers.CollectionFromSplit(split, "from")
	if err != dberrs.DB_NORM() {
		return err.Err
	}

	_, err = parsers.UrlFromSplit(split)
	if err != dberrs.DB_NORM() {
		return err.Err
	}
	return ""
}

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