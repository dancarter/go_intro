package main

import "errors"
import "fmt"

// In Go, it's idiomatic to communicate errors via an explicit,
// separate return value. This contrasts with the exceptions
// used in languages like Java and Ruby and the overloaded
// single result/error value sometimes used in C. Go's approach
// makes it easy to see which functions return errors and to handle
// them using the same language constructs employed for any other,
// non-error tasks

// Errors are the last return value and have type error,
// a built-in interface
func f1(arg int) (int, error) {
  if arg == 42 {
    return -1, errors.New("can't work with 42")
  }
  // A nil value in the error position indiciates there was no error
  return arg + 3, nil
}

// Custom types can be used as errors by implementing the Error()
// method on them
type argError struct {
  arg int
  prob string
}
func (e *argError) Error() string {
  return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// In this case we use &argError syntax to build a new struct,
// supplying values for the two fields arg and prob
func f2(arg int) (int, error) {
  if arg == 42 {
    return -1, &argError{arg, "can't work with 42"}
  }
  return arg + 3, nil
}

func main() {
  // The loops test out the error-returning functions
  // Inline error checking on the if line is a common
  // idiom in Go
  for _, i := range []int{7, 42} {
    if r, e := f1(i); e != nil {
      fmt.Println("f1 failed:", e)
    } else {
      fmt.Println("f1 worked:", r)
    }
  }
  for _, i := range []int{7, 42} {
    if r, e := f2(i); e != nil {
      fmt.Println("f2 failed:", e)
    } else {
      fmt.Println("f2 worked:", r)
    }
  }

  // If you want to programmatically use the data in a custom
  // error, you'll need to get the error as an instance of the
  // custom errortype via type assetion
  _, e := f2(42)
  if ae, ok := e.(*argError); ok {
    fmt.Println(ae.arg)
    fmt.Println(ae.prob)
  }
}
