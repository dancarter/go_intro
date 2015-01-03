package main

import s "strings"
import "fmt"

// The standard library's 'strings' package provides many
// useful string-related functions.

// Alias fmt.Println to a shorter name
var p = fmt.Println

func main() {
  // A sampling of functions available in 'strings'. These
  // are all methods from the package, not methods on the
  // string object itself, meaning we need to pass the
  // string in as the first argument
  p("Contains: ", s.Contains("test", "es"))
  p("Count: ", s.Count("test", "t"))
  p("HasPrefix: ", s.HasPrefix("test", "te"))
  p("HasSuffix: ", s.HasSuffix("test", "st"))
  p("Index: ", s.Index("test", "e"))
  p("Join: ", s.Join([]string{"a", "b"}, "-"))
  p("Repeat: ", s.Repeat("a", 5))
  p("Replace: ", s.Replace("foo", "o", "0", -1))
  p("Replace: ", s.Replace("foo", "o", "0", 1))
  p("ToLower: ", s.ToLower("TEST"))
  p("ToUpper: ", s.ToUpper("test"))
}
