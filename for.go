package main

import "fmt"

func main() {
  // basic for loop with single condition
  i := 1
  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }

  // initial/condition/after foor loop
  for j := 7; j<= 9; j++ {
    fmt.Println(j)
  }

  // for with no condition loops until you break out or
  // return from enclosing function
  for {
    fmt.Println("loop")
    break
  }
}
