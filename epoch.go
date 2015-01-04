package main

import "fmt"
import "time"

// A common requirement in programs is getting the number
// of seconds, milliseconds, or nanoseconds since the Unix
// epoch.

func main() {
  p := fmt.Println

  // Use time.Now with Unix or UnixNano to get elapsed time
  // since the Unix epoch in seconds or nanoseconds,
  // respectively
  now := time.Now()
  secs := now.Unix()
  nanos := now.UnixNano()
  p(now)

  // Note that there is no UnixMilli, so to get the
  // milliseconds since epoch you'll need to manually
  // divide from nanoseconds
  millis := nanos / 1000000
  p(secs)
  p(millis)
  p(nanos)

  // You can also convert integer seconds or nanoseconds
  // since the epoch into the corresponding time
  p(time.Unix(secs,0))
  p(time.Unix(0,nanos))
}
