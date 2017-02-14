  package main
  import (
    "fmt"
    "os"
  )
  
  func main(){
    if err := doError(); err != nil {
      fmt.Println("water", err)
      os.Exit(1)
    }
    fmt.Println("the end") 
  }

 func doError() error {
   msg := "nothing at all"
   return fmt.Errorf("fire %s", msg)
 }
