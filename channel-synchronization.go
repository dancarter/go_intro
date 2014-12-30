package main

import "fmt"
import "time"

// We can use channels to synchronize execution across goroutines


// This function will be run in a goroutine. The done channel will
// be used to notify another goroutine that this function is done
func worker(done chan bool) {
  fmt.Print("working...")
  time.Sleep(time.Second)
  fmt.Println("done")

  // We send a value to done to notify we're finished
  done <- true
}

func main() {
  // Start a worker goroutine giving it a channel to notify on
  done := make(chan bool, 1)
  go worker(done)

  // Block until we receive a notification from the worker
  <-done
  // If this line were removed, the program would exit before
  // the worker even started
}
