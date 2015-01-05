package main

import "fmt"
import "os"

// Use os.Exit to immediately exit with a given status

func main() {
  // defers will not be run when using os.Exit, so this
  // fmt.Println will never be called
  defer fmt.Println("!")

  // Exit with status 3
  os.Exit(3)

  // Note that unlike eg C, Go does not use an integer return
  // value from main to indicate exit status. If you'd like to
  // exit with a non-zero status, you should use os.Exit
}
