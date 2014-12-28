package main

import "fmt"

// Go's structs are typed collections of fields

// This person struct has name and age fields
type person struct {
  name string
  age int
}

func main() {
  // This syntax creates a new struct
  fmt.Println(person{"Bob", 20})

  // You can name fields when initializing
  fmt.Println(person{name: "Alice", age: 30})

  // Omitted fields will be zero-valued
  fmt.Println(person{name: "Fred"})

  // An & prefix yields a pointer to the struct
  fmt.Println(&person{name: "Ann", age: 40})

  // Access struct fields with a dot
  s := person{name: "Sean", age: 50}
  fmt.Println(s.name)

  sp := &s
  fmt.Println(sp.age)

  sp.age = 51
  fmt.Println(sp.age)
}
