package main

import (
	"fmt"
)

type Product struct{
  name string
  price int
  discount int
}

func (product Product) discountDiffrence() string{
  return fmt.Sprintf("your discount diffrence is ==> %v", product.price - (product.price * product.discount / 100))
} 

func run() string{
  rob := Product{name:"rob goje", price:2500, discount:10}
	return rob.discountDiffrence()
}

func main() {
  fmt.Println(run())
}
