package main

import "fmt"

// Channels are the pipes that connect concurrent goroutines
// You can send values into channels from one goroutine,
// and receive them in another

func main() {
  // A new channel is created with 'make(chan val-type)'
  // Channels are typed by the values they convey
  messages := make(chan string)

  // A value is sent into a channel using the syntax:
  // channel <- value
  go func() { messages <- "ping" }()

  // A value is retrieved from a channel with the syntax:
  // <-channel
  msg := <-messages
  fmt.Println(msg)

  //By default, send and receive block until both the sender
  // and the receiver are ready. This property allows the program
  // to wait at the end for the "ping" message
}
