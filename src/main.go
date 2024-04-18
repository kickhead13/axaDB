package main

import (
 "os"
 "axaDB/src/parsers"
 "axaDB/src/axa_init"
 "fmt"
)

func main() {
    var args []string = os.Args[1:]
    if len(args) == 0 {
        fmt.Println(parsers.AxaHelp())        
    } else {
        switch args[0] {
        case "init":
            axa_init.Init(args[1:])
        }
    }
}