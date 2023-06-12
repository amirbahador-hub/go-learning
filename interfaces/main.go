package main

import (
  "fmt"
)

type Animal interface {
  walk()
}

type Cat struct {
  name string 
}

func (cat Cat) walk() {
  fmt.Println("Cat is walking")
}

func run(animal Animal){
  animal.walk()
}

func main(){
  pashmak := Cat{"pashmak"}
  run(pashmak)
}
