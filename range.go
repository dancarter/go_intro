package main

import "fmt"

func main() {
  // a range iterates over elements in a variety of data structures

  // using range to sum the numbers in a slice
  // an array would work the same way
  nums := []int{2, 3, 4}
  sum := 0
  for _, num := range nums {
    sum += num
  }
  fmt.Println("sum:", sum)

  // using range on an array or slice provides both the index
  // and the value for each entry. We didn't need the index before,
  // so we ignored it with the 'blank identifier' _
  // Sometimes we will want to use the range
  for i, num := range nums {
    if num == 3 {
      fmt.Println("index:", i)
    }
  }

  // range on map iterates over key/value pairs
  kvs := map[string]string{"a": "apple", "b": "banana"}
  for k, v := range kvs {
    fmt.Printf("%s -> %s\n", k, v)
  }

  // range on strings iterates over Unicode code points
  // The first value is the starting byte idnex of the 'rune'
  // and the second the 'rune' itself
  for i, c := range "go" {
    fmt.Println(i, c)
  }
}
