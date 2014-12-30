package main

import "fmt"
import "time"

// Here is a worker several of which will be run concurrently.
// These workers receive work on te jobs channel and send the
// results to on the reults channel. We'll sleep a second per
// job to simulate an expensive task
func worker(id int, jobs <- chan int, results chan<- int) {
  for j:= range jobs {
    fmt.Println("worker", id, "processing job", j)
    time.Sleep(time.Second)
    results <- j * 2
  }
}

func main() {
  // To use a pool of workers, they need to be sent work and
  // have their results collected. Two channels will be made
  // for this purpose
  jobs := make(chan int, 100)
  results := make(chan int, 100)

  // Three workers are started up and blocked because there
  // are not yet any jobs
  for w := 1; w <= 3; w++ {
    go worker(w, jobs, results)
  }

  // Sending nine jobs and then closing the jobs channel to
  // indicate there is no more work
  for j := 1; j <= 9; j++ {
    jobs <- j
  }
  close(jobs)

  // Finally, collect the results
  for a := 1; a <= 9; a++ {
    <-results
  }

  // The running program showsthe 9 jobs being executed by various
  // workers. The program only takes about 3 seconds despite doing
  // 9 seconds total work because there are 3 workers operating
  // concurrently.
}
