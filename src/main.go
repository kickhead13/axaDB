package main

import (
 "os"
 "fmt"
 //"reflect"
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

    
  str := "7DzgywhxPrwwfaXDyO0nHU38Ym2GdjjtGGQ6EQ+fro6Iqz0g9+ZNC0dqoIGlyVobDouliIrysfr1Bq9eCZZIOw=="
  dstr, err := axa_security.DecryptData(str)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(string(dstr))
  var args []string = os.Args[1:]
  if program := parsers.BoolParse(args, []string{"start", "init", "connect", "halt"}); !program {
    if help := parsers.BoolParse(args, []string{"-h", "--help"}); help {
      fmt.Println(parsers.AxaHelp())
      return;
    }
  }
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
