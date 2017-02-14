  package main
  import (
    "fmt"
  )

  type Animal struct {
    Name string
    Age int
  }
 
  type Cat struct {
   Animal // 埋め込み
   ServantName string
  }
  
 func main(){
   // c := Cat{} 
   // c.Name = "siro"
   // c := Cat{Animal: Animal{Name:"hoge", 5}, ServantName:"Jiro"}
   // c.Name = "Taro"
   a1 := &Animal{"Hoge", 5}
   // a2 := new(Animal{"Hoge", 5})
   // a.Name = "Hoge"
   // a.Age = 5
   fmt.Println(a1.Name) // => siro , Jiro, Hoge, Hoge
 }
