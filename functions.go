package main

import "fmt"

// This function takes two ints and returns an int
func plus(a int, b int) int {
  // Go requires explicit returns and won't automatically return
  // the value of the last expression executed
  return a + b
}

func main() {
  // Call a function with 'name(args)'
  res := plus(1, 2)
  fmt.Println("1 + 2 =", res)
}
