package main

import (
  "fmt"
  "time"
  "sync"
)

func sleeper(name string, a time.Duration,wg *sync.WaitGroup){
  defer wg.Done()
  fmt.Printf("Start %v \n", name)
  time.Sleep(a * time.Millisecond)
  fmt.Printf("Done %v \n", name)
}

//func sequential(){
//  sleeper("Downloading the image", 5000)
//  sleeper("Convert the image", 3000)
//  sleeper("Saving the image into the database", 4000)
//}

func main(){
  wg := &sync.WaitGroup{}

  wg.Add(3)
  go sleeper("Downloading the image", 5000, wg)
  go sleeper("Convert the image", 3000, wg)
  go sleeper("Saving the image into the database", 4000, wg)
  wg.Wait()
  fmt.Println("All Done")
  //sleeper("THe END", 5000)
}

