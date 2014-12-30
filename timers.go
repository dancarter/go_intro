package main

import "time"
import "fmt"

// Go's built in timer and ticker features make executing
// Go code at some point in the future, or repeatedly at some
// interval, easy.

func main() {
  // Timers represent a single event in the future.
  // You tell the timer how long to want to wait, and it
  // provides a channel that will be notified at that time.
  // This timer waits 2 seconds
  timer1 := time.NewTimer(time.Second * 2)

  // The '<-timer1.C' blocks the timer's channel C until it
  // sends a value indicating that the timer expired
  <-timer1.C
  fmt.Println("Timer 1 expired")

  // If you just want to wait, time.Sleep can be used. One
  // reason a timer may be useful is that it can be cancelled
  // before the timer expires.
  timer2 := time.NewTimer(time.Second)
  go func() {
    <-timer2.C
    fmt.Println("Timer 2 expired")
  }()
  stop2 := timer2.Stop()
  if stop2 {
    fmt.Println("Timer 2 stopped")
  }

  // The first timer expires ~2 seconds after the program starts,
  // but the second will be stopped before it expires
}
