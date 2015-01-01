package main

import "sort"
import "fmt"

// Sometimes, there will be a need to sort a collection by something
// other than its natural order. For example, suppose we wanted to
// sort strings by their length instead of alphabetically.

// In order to sort by a custom function in Go, we need a
// corresponding type. Here, a ByLength type is created that is just
// an alias for the builtin []string type
type ByLength []string

// We implement sort.Interface - Len, Less, and Swap - on our type
// so we can use the sort packages' generic Sort function. Len and
// Swap will usally be similar across types and Less will hold the
// actual custom sorting logic. In our case, we want to sort in order
// of increasing string length, so we use len(s[i]) and len(s[j])
func (s ByLength) Len() int {
  return len(s)
}
func (s ByLength) Swap(i, j int) {
  s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
  return len(s[i]) < len(s[j])
}

func main() {
  // With all of this in place, we can now implement our custom sort
  // by casting the original fruits slice to ByLength, and then use
  // sort.Sort on that typed slice
  fruits := []string{"peach", "banana", "kiwi"}
  sort.Sort(ByLength(fruits))
  fmt.Println(fruits)
}
