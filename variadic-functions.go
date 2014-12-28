package main

import "fmt"

// Variadic functions can be called with any number of
// trailing arguments. For example, 'fmt.Println' is a
// common variadic function

// This function takes an arbitrary number of ints
func sum(nums ...int) {
  fmt.Print(nums, " ")
  total := 0
  for _, num := range nums {
    total += num
  }
  fmt.Println(total)
}

func main() {
  // Variadic functions are called the usual way
  // with individual arguments
  sum(1, 2)
  sum(1, 2, 3)

  // A slice can be passed in using func(slice...)
  nums := []int{1, 2, 3, 4}
  sum(nums...)
}
