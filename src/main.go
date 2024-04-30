package main

import (
 "os"
 "fmt"
 "reflect"
)

import (
 "axaDB/src/parsers"
 "axaDB/src/axa_init"
 "axaDB/src/dberrs"
 "axaDB/src/axa_connect"
 "axaDB/src/axa_server"
 "axaDB/src/axa_security"
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
        case "connect":
            err := axa_connect.Connect(args[1:])
            if err != dberrs.DB_NORM() {
                fmt.Println(err.Err)
            }
        case "start":
            err := axa_server.Start(args[1:])
            if err != dberrs.DB_NORM() {
                fmt.Println(err.Err)
            }
        }
    
    }
}