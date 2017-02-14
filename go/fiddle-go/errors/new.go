  package main

  import (
    "errors"
    "os"
    "fmt"
  )

  func main(){
    if err := doErrors(); err != nil {
      fmt.Println("err", err)
      os.Exit(1) 
    }
    fmt.Println("the end") 
  }

  func doErrors() error {
    return errors.New("nothing at all") 
  }
