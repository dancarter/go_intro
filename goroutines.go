package main

import "fmt"

// A goroutine is a lightweight thread of execution

func f(from string) {
  for i := 0; i < 3; i++ {
    fmt.Println(from, ":", i)
  }
}

func main() {
  // Calling f(s) the usual way it runs synchronously
  f("direct")

  // To invoke the function in a goroutine, we add 'go'
  // in front of the function call. This goroutine will
  // execute concurrently with the calling one
  go f("goroutine")

  // goroutines can also be started for anonymous functions
  go func(msg string) {
    fmt.Println(msg)
  }("going")

  // The two goroutines that have been started will run
  // asynchronously in seperate goroutines, so execution
  // falls through to here. 'Scanln' requires a key press
  // before the program will exit
  var input string
  fmt.Scanln(&input)
  fmt.Println("done")

  // When the program is run, the output of the calling
  // block will be seen first, then the interleaved output
  // of the two goroutines.
}
