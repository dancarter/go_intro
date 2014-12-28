package main

import "fmt"
import "time"

func main() {
  // basic switch
  i := 2
  fmt.Println("write ", i, " as ")
  switch i {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
  case 3:
    fmt.Println("three")
  }

  // commas can seperate multiple expressions in
  // the same case statement
  // a "default" case is run when no other case matches
  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("it's the weekend")
  default:
    fmt.Println("it's a weekday")
  }

  // a switch with no expression is an alternative
  //  way to express if/else logic
  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("it's before noon")
  default:
    fmt.Println("it's after noon")
  }
}
