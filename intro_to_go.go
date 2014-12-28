package main

import (
  "fmt"
)

// capitalizing the first letter of the function makes it
// available outside if this where to be imported. Basically:
// lowercase = private function
// uppercase = public function
func Greet(name string) {
  fmt.Println("Hello, " + name)
}

func main() {
  // Long form variable declaration  - var name string = "Daniel"
  // Short form variable declaration - name := "Daniel"

  names := []string {
    "Daniel",
    "Mark",
    "Rachel",
    "Dylan",
    "Leo",
  }

  // _     - index, not being used so _ throws it away.
  //         Can be set to a variable to be used if needed
  // n     - the value at the current index in the array
  // range - takes the array names and passes back each
  //         value with its index
  for _, n := range names {
    Greet(n)
  }
}
