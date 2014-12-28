package main

import "fmt"

func main() {
  // Maps are Go's built-in associative data type
  // (also known as hashes or dicts in other languages)

  //An empty map is created with make:
  // make(map[key-type]val-type)
  m := make(map[string]int)

  // Set key/value pairs with name[key] = value
  m["k1"] = 7
  m["k2"] = 13

  // Printing with Println shows all key/value pairs
  fmt.Println("map:", m)

  // Get a value with name[key]
  v1 := m["k1"]
  fmt.Println("v1:", v1)

  // 'len' will return the number of key/value pairs
  fmt.Println("len:", len(m))

  // 'delete' will remove a key/value pair
  delete(m, "k2")
  fmt.Println("map:", m)

  // When getting a value, a map returns and optional
  // second return value which indicates if a key is present
  // This is used to disambiguate missing keys and keys with
  // a value of 0 or ""
  _, prs := m["k2"]
  fmt.Println("prs:", prs)

  // A new map can be declared and initialized on one line
  n := map[string]int{ "foo": 1, "bar": 2}
  fmt.Println("map:", n)
}
