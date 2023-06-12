package main

import (
  "fmt"
  "time"
  "math/rand"
  "strings"
)

func downloading_the_image(chn chan string, link string){

   // Generate random number between 0-9
   rand.Seed(time.Now().UnixNano())
   randNum := rand.Intn(10)

   // downloading the image
   fmt.Printf("Start downloading the %v, it might take %v secounds \n", link, randNum)
   time.Sleep(time.Duration(randNum * 1000) * time.Millisecond)
   fmt.Printf("Finish downloading the %v \n", link)

   splited_link := strings.Split(link, "/")
   dlink := fmt.Sprintf("https://superzz.ir/%v",splited_link[len(splited_link)-1])
   chn <- dlink
}

func main(){
  chn := make(chan string, 3)
  go downloading_the_image(chn, "https://google.com/hmid.png")
  go downloading_the_image(chn, "https://digikala.com/laptop.png")
  go downloading_the_image(chn, "https://snapp.com/taxi.svg")

  dlink1 := <- chn
  fmt.Println(dlink1)

  dlink2 := <- chn
  fmt.Println(dlink2)

  dlink3 := <- chn
  fmt.Println(dlink3)

}
