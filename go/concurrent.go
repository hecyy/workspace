package main

import (
	"fmt"
  "time"
)

func boring (msg string) (c chan string) {
  c = make(chan string)
  go func() {
    for i := 0 ;;i++ {
      c <- fmt.Sprintf("%s %d", msg, i)
      time.Sleep(5*time.Second)
    }
  }()
  return
}

func main() {
  c := boring("this is so boring")
  for {
    select {
      case s :=  <-c:
        fmt.Println(s)
      case s := <-time.After(1*time.Second):
        fmt.Println("Too late!!", s)
        return
    }
  }
}
