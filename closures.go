package main

import "fmt"

// Go supports 'anonymous functions', which can form
// 'closures'
// Anonymous functions are useful when you want to
// define a function inline without naming it

// This function returns another function which is
// defined anonymously in the body of intSeq
// The returned function closes over the variabl i
func intSeq() func() int {
  i := 0
  return func() int {
    i += 1
    return i
  }
}

func main() {
  // intSeq is called and the returned function is assigned
  // to a variable. This function captures its own i values,
  // which will be updated each time we call nextInt
  nextInt := intSeq()

  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())

  // A new function created with intSeq will have a new i value
  newInts := intSeq()
  fmt.Println(newInts())
}
