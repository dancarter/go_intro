package main

import (
  "bytes"
  "fmt"
  "regexp"
)

// Go offers built-in support for regular expressions

func main() {
  p := fmt.Println
  // This tests if a pattern matches a string
  match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
  p(match)

  // For other regexp tasks, you'll need to 'Compile' an
  // optimized Regexp struct
  r, _ := regexp.Compile("p([a-z]+)ch")

  // Many methods are available on Regexp structs
  p(r.MatchString("peach"))

  //Finds the match for the regexp
  p(r.FindString("peach punch"))

  // Finds the first match but returns the start and end indexes
  // for the match instead of the matching text
  p(r.FindStringIndex("peach punch"))

  // The Submatch variants include information about both the
  // whole-pattern matchs and the submatches within those
  // matches. For example, this will return information for both
  // 'p([a-z]+)ch' and '([a-z]+)'
  p(r.FindStringSubmatch("peach punch"))

  // Similarly, this will return information about the indexes
  // of matches and submatches
  p(r.FindStringSubmatchIndex("peach punch"))

  // The All variants of these functions apply to all matches
  // in the input not just the first.
  p(r.FindAllString("peach punch pinch", -1))

  // The All variants are available for the other functions
  // as well
  p(r.FindAllStringIndex("peach punch pinch", -1))
  p(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

  // Providing a non-negative integer as the second argument
  // to these functions will limit the number of matches
  p(r.FindAllString("peach punch pinch", 2))

  // Our examples above had string arguments and used names
  // like MatchString. We can also provide []byte arguments
  // and drop the String from the function name
  p(r.Match([]byte("peach")))

  // When creating constants with regular expressions, you
  // can use the MustCompile variation of Compile. A plain
  // Compile won't work for constants because it has 2
  // return values
  r = regexp.MustCompile("p([a-z]+)ch")
  p(r)

  // The regexp package can also be used to replace subsets
  // of strings with other values.
  p(r.ReplaceAllString("a peach", "<fruit>"))

  // The Func variant allows you to transform matched text
  // with a given function.
  in := []byte("a peach")
  out := r.ReplaceAllFunc(in, bytes.ToUpper)
  p(string(out))
}
