package main

import "strconv"
import "fmt"

// Parsing numbers from strings is a basic but common task
// in many programs

func main() {
  p := fmt.Println

  // The built-in package strconv provides the number parsing
  f, _ := strconv.ParseFloat("1.234", 64)
  p(f)

  // For ParseInt, the 0 means infer the base from the string.
  // 64 requires that the result fit in 64 bits.
  i, _ := strconv.ParseInt("123", 0, 64)
  p(i)

  // ParseInt will recognize hex-formatted numbers
  d, _ := strconv.ParseInt("0x1c8", 0, 64)
  p(d)

  // A ParseUint is also available
  u, _ := strconv.ParseUint("789", 0, 64)
  p(u)

  // Atoi is a convenience function for basic base-10 int
  // parsing
  k, _ := strconv.Atoi("135")
  p(k)

  // Parse functions return an error on bad input
  _, e := strconv.Atoi("wat")
  p(e)
}
