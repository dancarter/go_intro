package main

import "fmt"

// Go has built-in support for multiple return values
// This is used often in idiomatic Go
// For example, both result and error values from a function

// (int, int) in the function signature shows the function
// returns 2 integer values
func vals() (int, int) {
  return 3, 7
}

func main() {
  // We use the two different return values with 'multiple assignment'
  a, b := vals()
  fmt.Println(a)
  fmt.Println(b)

  // If you only want a subset of returned values,
  // the 'blank identifier' can be used
  _, c := vals()
  fmt.Println(c)
}
