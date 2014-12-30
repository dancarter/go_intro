package main

import "time"
import "fmt"

// Timers are for when you want to do something once in the
// future, but tickers are for when you want to do something
// repeatedly at regular intervals.

func main() {
  // Tickers use a similar mechanism to timers: a channel
  // that is sent values. Using a range on the channel, we
  // can iterate over the values as they arrive every 500ms
  ticker := time.NewTicker(time.Millisecond * 500)
  go func() {
    for t := range ticker.C {
      fmt.Println("Tick at", t)
    }
  }()

  time.Sleep(time.Millisecond * 1500)

  // Tickers can be stopped like timers. Once a ticker is
  // stopped, it won't receive any more values on its channel
  ticker.Stop()
  fmt.Println("Ticker Stopped")

  //When the program is run, the ticker should tick 3 times
}
