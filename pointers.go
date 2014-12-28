package main

import "fmt"

// Go supports pointers, allowing you to pass references
// to values and records within your program

// This function has an int parameter, so arguments are
// passed to it by value. 'zeroval' will get a copy of ival
// distinct from the one in the calling function
func zeroval(ival int) {
  ival = 0
}

// This function has an *int parameter, meaning it receives
// a pointer to an int. The *iptr code in the function
// 'dereferences' the pointer from its memory address to
// the current value at that address. Assigning a value
// to a dereferenced pointer changes the value at the
// referenced address
func zeroptr(iptr *int) {
  *iptr = 0
}

func main() {
  i := 1
  fmt.Println("initial:", i)

  zeroval(i)
  fmt.Println("zeroval:", i)

  //The &i syntax gives the memory address of i
  // ie a pointer to i
  // Pointers can be printed too
  zeroptr(&i)
  fmt.Println("zeroptr:", i)

  fmt.Println("pointer:", &i)

  // zeroval doesn't change i in main, but zeroptr does
}
