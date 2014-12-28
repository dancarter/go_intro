package main

import "fmt"

func main() {
  // slices are typed only by type of elements,
  // and not size like arrays. To create an empty
  // slices with non-zero length, 'make' is used
  s := make([]string, 3)

  // set/get work the same as arrays
  s[0] = "a"
  s[1] = "b"
  s[2] = "c"
  fmt.Println("set:", s)
  fmt.Println("get:", s[2])

  // len returns the length of a slice
  fmt.Println("len:", len(s))

  // 'append' can be used to returns a slice
  // containing one of more new values.
  // returns a new slice value
  s = append(s, "d")
  s = append(s, "e", "f")
  fmt.Println("apd:", s)

  // Slices can also be copied with 'copy'
  // The contents of s are copied into c
  c := make([]string, len(s))
  copy(c,s)
  fmt.Println("cpy:", c)

  // The 'slice' operator with the syntax
  // 'slice[low:high]' gets elements from
  // low to high - 1
  l := s[2:5]
  fmt.Println("sl1:", l)

  // Omitting low value starts from 0
  l = s[:5]
  fmt.Println("sl2:", l)

  // Omitting high value ends at the end of the slice
  l = s[2:]
  fmt.Println("sl3:", l)

  // A slice and be delcared and initialized in one line
  t := []string{"g", "h", "i"}
  fmt.Println("dcl:", t)

  // Slices can be multi-dimensional
  // unlike arrays, the length of inner-slices can vary
  twoD := make([][]int, 3)
  for i := 0; i < 3; i++ {
    innerLen := i + 1
    twoD[i] = make([]int, innerLen)
    for j := 0; j < innerLen; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2d:", twoD)
}
