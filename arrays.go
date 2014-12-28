package main

import "fmt"

func main() {
  // a is an array that will hold exactly 5 ints
  // the type of elements and the length are both
  // part of the array's type
  // by default an array is zero-valued
  var a [5]int
  fmt.Println("empty: ", a)

  //values can be set at an index using the
  // array[index] = value syntax, and get a value
  // with array[index]
  a[4] = 100
  fmt.Println("set:", a)
  fmt.Println("get:", a[4])

  // The built in len method returns an array's length
  fmt.Println("length:", len(a))

  // this syntax declares and initializes
  // an array in one line
  b := [5]int{1, 2, 3, 4, 5}
  fmt.Println("dcl:", b)

  // you can compose types to build multi-dimensional
  // data structures
  var twoD[2][3]int
  for i := 0; i < 2; i++ {
    for j := 0; j < 3; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2d: ", twoD)
}
