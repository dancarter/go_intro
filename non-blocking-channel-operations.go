package main

import "fmt"

// Basic sends and receives on channels are blocking.
// However, we can use 'select' with a 'default' clause
// to implement 'non-blocking' sends, receives, and
// even non-blocking multi-way selects

func main() {
  messages := make(chan string)
  signals := make(chan bool)

  // This is a non-blocking receive.
  // If a value is available on messages, then the
  // select will take the <-messages case with that
  // value. If not, it will immediately take the default
  select {
  case msg := <-messages:
    fmt.Println("received message", msg)
  default:
    fmt.Println("no received message")
  }

  // A non-blocking send works similarly
  msg := "hi"
  select {
  case messages <- msg:
    fmt.Println("sent message", msg)
  default:
    fmt.Println("no message sent")
  }

  // We can use multiple cases above the default to
  // implement a multi-way non-blocking select.
  // Here non-blocking receives are attempted on
  // both signals and messages
  select {
  case msg := <-messages:
    fmt.Println("received message", msg)
  case sig := <-signals:
    fmt.Println("received signal", sig)
  default:
    fmt.Println("no activity")
  }
}
