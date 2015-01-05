package main

import (
  "os"
  "strings"
  "fmt"
)

// Environment variables are a universal mechanism for
// conveying configuration information to Unix programs.

func main() {
  // To set a key/value pair, use os.Setenv. To get a value
  // for a key, use os.Getenv. This will return an empty
  // string if the key isn't present in the environment.
  os.Setenv("FOO", "1")
  fmt.Println("FOO:", os.Getenv("FOO"))
  fmt.Println("BAR:", os.Getenv("BAR"))

  // Use os.Environ to list all key/value pairs in the
  // environment. This returns a slice of strings in the
  // form KEY=value. You can string.Split them to get the
  // key and value.
  fmt.Println()
  for _, e := range os.Environ() {
    pair := strings.Split(e, "=")
    // Print the keys
    fmt.Println(pair[0])
  }
}
