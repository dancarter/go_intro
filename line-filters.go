package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

// A line filter is a common type of program that reads
// input on stdin, processes it ,and then prints some
// derived result to stdout. grep and sed are common
// line filters.

// This is an example of a line filter in Go that writes
// a capitalized version of all input text. You can use
// this pattern to write other Go line filters

func main() {
  // Wrapping the unbuffered os.Stdin with a buffered
  // scanner gives us a convenient Scan method that advances
  // the scanner to the next token; which is the next line
  // in the default scanner.
  scanner := bufio.NewScanner(os.Stdin)

  // Text returns the current token, here the next line,
  // from the input
  for scanner.Scan() {
    ucl := strings.ToUpper(scanner.Text())

    // Write out the uppercased line
    fmt.Println(ucl)
  }

  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "error:", err)
    os.Exit(1)
  }
}
