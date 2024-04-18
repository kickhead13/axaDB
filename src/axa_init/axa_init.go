package axa_init

import (
	"fmt"
	"axaDB/src/parsers"
	"os"
)

func Init(args []string) {
	if len(args) == 0 {
		fmt.Println("????")
	} else {
		at, _ := parsers.InitParse(args, []string{"--at", "-a"}[:])
		_ = os.Mkdir(at, 0755)
	}
}

