package main

import "time"
import "fmt"

// Rate limiting is an important mechanism for controlling
// resource utilization and maintaining quality of service.
// Go elegantly supports rate limiting with goroutines,
// channels, and tickers.

func main() {
  // First, basic rate limiting. Suppose we want to limit
  // the handling of incomin requests. We'll serve these
  // requests off a channel of the same name.
  requests := make(chan int, 5)
  for i:= 1; i <= 5; i++ {
    requests <- i
  }
  close(requests)

  // The 'limiter' channel will receive a value every 200ms
  // This is the regulator in our rate limiting scheme
  limiter := time.Tick(time.Millisecond * 200)

  // By blocking on a receive from the limiter channel before
  // serving each request, we limit ourselves to 1 request
  // every 200 milliseconds
  for req := range requests {
    <-limiter
    fmt.Println("request", req, time.Now())
  }

  // We may want to allow short bursts of requests in our rate
  // limiting scheme while preserving the overall rate limit. We
  // can accomplish this by buffering our limiter channle. This
  // burstyLimiter channel will allow bursts of up to 3 events
  burstyLimiter := make(chan time.Time, 3)

  // Fill the channel to represent allowed bursting
  for i:= 0; i < 3; i++ {
    burstyLimiter <- time.Now()
  }

  // Every 200ms we'll try to add a new value to burstyLimiter,
  // up to its limit of 3
  go func() {
    for t:= range time.Tick(time.Millisecond * 200) {
      burstyLimiter <- t
    }
  }()

  // Now simulate 5 more incomin requests. The first 3 of these
  // will benefit from the burst capability of burstyLimiter
  burstyRequests := make(chan int, 5)
  for i := 1; i <= 5; i++ {
    burstyRequests <- i
  }
  close(burstyRequests)
  for req := range burstyRequests {
    <-burstyLimiter
    fmt.Println("request", req, time.Now())
  }

  // Running the program, the first batch of requests are handled
  // once every 200 milliseconds

  // The second batch of requests we serve the first 3 immediately
  // because of the burstable rate limiting, then serve the remaining
  // 2 with _200ms delays each
}
