package main

import "os"

// A panic typically means something went unexpectedly wrong.
// Mostly we use it to fail fast on errors that shouldn't
// occur during normal operation, or that we aren't prepared
// to handle gracefully

func main() {
  panic("a problem")

  // A commun use of panic is to abort if a function returns
  // an error value that we don't know how to (or want to)
  // handle. Here is an example of panicking if we get an
  // unexpected error when creating a new file
  _, err := os.Create("/tmp/file")
  if err != nil {
    panic(err)
  }

 // Not far from some languages which use exceptions for
 // handling of many errors, in Go it is idiomatic to use
 // error-indicating return values wherever possible
}
