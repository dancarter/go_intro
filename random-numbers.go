package main

import "fmt"
import "math/rand"

// Go's math/rand package provides pseudorandom number
// generation.

func main() {
  p := fmt.Print

  // For examples, rand.Intn returns a random int n,
  // 0 <= n < 100
  p(rand.Intn(100), ",")
  p(rand.Intn(100))
  fmt.Println()

  // rand.FLoat64 returns a float64 f, 0.0 <= f < 1.0
  fmt.Println(rand.Float64())

  // this can be used to generate random floats in other
  // ranges, for example 5.0 <= f < 10.0
  p((rand.Float64()*5)+5, ",")
  p((rand.Float64() * 5) + 5)
  fmt.Println()

  // To make the pseudorandom generator deterministic,
  // give it a well-known seed.
  s1 := rand.NewSource(42)
  r1 := rand.New(s1)

  // Call the resulting rand.Sourcejust like the functions
  // on the rand package
  p(r1.Intn(100), ",")
  p(r1.Intn(100))
  fmt.Println()

  // If you seed a source with the same number, it produces
  // the same sequence of random numbers.
  s2 := rand.NewSource(42)
  r2 := rand.New(s2)
  p(r2.Intn(100), ",")
  p(r2.Intn(100))
  fmt.Println()
}
