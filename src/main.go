package main

import (
 "os"
 "axaDB/src/parsers"
 "axaDB/src/axa_init"
 "fmt"
 "axaDB/src/dberrs"
)

func main() {
    var args []string = os.Args[1:]
    if len(args) == 0 {
        fmt.Println(parsers.AxaHelp())        
    } else {
        switch args[0] {
        case "init":
            err := axa_init.Init(args[1:])
            if err != dberrs.DB_NORM() {
                fmt.Println(err.Err)
            }
        }
    }
}