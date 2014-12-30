package main

import "fmt"
import "time"

// Go's 'select' lets you wait on multiple channel operations
// Combining goroutines and channels with select is a
// powerful feature of Go

func main() {
  // We will select across two channels
  c1 := make(chan string)
  c2 := make(chan string)

  // Each channel will receive a value after some amount of time
  go func() {
    time.Sleep(time.Second * 1)
    c1 <- "one"
  }()
  go func() {
    time.Sleep(time.Second * 2)
    c2 <- "two"
  }()

  // Use 'select' to await both of these values simultaneously
  // and print them out as they arrive
  for i := 0; i < 2; i++ {
    select {
    case msg1 := <-c1:
      fmt.Println("received", msg1)
    case msg2 := <-c2:
      fmt.Println("received", msg2)
    }
  }
  // When the program is run, the strings 'one' and then 'two' are
  // received. Since both Sleeps execute concurrently, execution time
  // is only ~2 seconds
}
