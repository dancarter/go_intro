package main

import "fmt"

// Closing a channel indicates no more values will be
// sent on it. This can be useful to communicate completion
// to the channel's receivers

func main() {
  // The jobs channel will be used to communicate work to be
  // done from the main() goroutineto a worker goroutine.
  // When there are no more jobs for the worker, we'll
  // close the jobs channel
  jobs := make(chan int, 5)
  done := make(chan bool)

  go func() {
    for {
      j, more := <- jobs
      if more {
        fmt.Println("received job", j)
      } else {
        fmt.Println("received all jobs")
        done <- true
        return
      }
    }
  }()

  // Send 3 jobs to the worker over the jobs channel and
  // then close the channel
  for j := 1; j <= 3; j++ {
    jobs <- j
    fmt.Println ("sent job", j)
  }
  close(jobs)
  fmt.Println("sent all jobs")

  // We wait for the worker using the synchronization approach
  <-done
}
