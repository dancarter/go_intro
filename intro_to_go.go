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

func GreetNames(names []string, suffix string) {
  // _     - index, not being used so _ throws it away.
  //         Can be set to a variable to be used if needed
  // n     - the value at the current index in the array
  // range - takes the array names and passes back each
  //         value with its index
  for _, n := range names {
    Greet(n + suffix)
  }
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


  // make keyword will create a new channel for us to use
  // here we are creating a channel that expects strings
  // to go back and forth
  comm := make(chan string)

  // go  - creates a new go routine to run the method in a
  //       new thread
  // <C> - concurrent loop
  // <M> - main loop
  go func(){ // anonymous function
    GreetNames(names, " <C> ")
    // <- is putting information into the comm channel
    comm <- "Finished greeting names concurrently..."
  }() //this runs function
  GreetNames(names, " <M> ")

  // Causes the program to wait for communication to come
  // out of the comm channel
  fmt.Println(<- comm)
}
