package main

import "fmt"

// We can use for and range to iterate over values
// received from a channel

func main() {
  // Iterating over two values in the queue channel
  queue := make(chan string, 2)
  queue <- "one"
  queue <- "two"
  close(queue)

  // This range iterates over each element as it is
  // received from the queue. Because we closed the channel
  // above, the iteration terminates after receiving the 2
  // elelments.
  for elem := range queue {
    fmt.Println(elem)
  }

  // This example also shows it is possible to close a
  // non-empty channel but still receive the remaining
  // values
}
