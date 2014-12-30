package main

import "fmt"
import "time"

// Timeouts are importnat for programs that connect
// to external resources or that otherwise need to
// bound execution time. Implementing timeouts in Go
// is easy and elegant thanks to channels and select

func main() {
  // Suppose we're executing an externa call that returns
  // its result on a channel 'c1' after 2 seconds
  c1 := make(chan string, 1)
  go func() {
    time.Sleep(time.Second * 2)
    c1 <- "result 1"
  }()

  // Here is a 'select' implementing a timeout
  // 'res := <-c1' awaits the result and
  // '<-Time.After' awaits a value to be sent
  // after the timeout of 1s. Since the 'select' proceeds
  // with the first receive that is ready, we'll take
  // the timeout case if the operation takes more
  // than the allowed 1s
  select {
  case res := <-c1:
    fmt.Println(res)
  case <-time.After(time.Second * 1):
    fmt.Println("timeout 1")
  }

  // If we allow a longer timeout of 3 seconds, then the receive
  // from c2 will succeed and print the result
  c2 := make(chan string, 1)
  go func() {
    time.Sleep(time.Second * 2)
    c2 <- "result 2"
  }()
  select {
  case res := <-c2:
    fmt.Println(res)
  case <-time.After(time.Second * 3):
    fmt.Println("timeout 2")
  }

  // When the program is run, the first operation times
  // out and the second succeeds

  // Using this 'select' timeout pattern requires
  // communicating results over channels. This is a good
  // idea in general because other important Go features
  // are based on channels and select
}
