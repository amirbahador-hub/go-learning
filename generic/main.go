package main

import (
  "fmt"
)


func run[T comparable](a T,b T){
  c := a == b
  fmt.Println(c)
}

func main(){
  run("Hamid","AmirBahador")
  run(2,2)
}
