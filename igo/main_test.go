package main

import "testing"

func TestRun(t *testing.T) {
// your discount diffrence is ==> 2250
  output := run()
  if output != "your discount diffrence is ==> 2250" {
    t.Errorf("run function output is WRONG")
  }
}
