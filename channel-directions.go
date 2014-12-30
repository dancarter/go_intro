package main

import "fmt"
// When using channels as function parameters, you can
// specify if a channel is meant to only send or receive
// This increases the type-safety of the program

// The 'ping' function accepts a channel for sending values
// If would be a compile-time error to try to receive
func ping(pings chan<- string, msg string){
  pings <- msg
}

// The 'pong' function accepts one channel for receives,
// and a second for sends
func pong(pings <-chan string, pongs chan<- string) {
  msg := <-pings
  pongs <- msg
}

func main() {
  pings := make(chan string, 1)
  pongs := make(chan string, 1)
  ping(pings, "passed message")
  pong(pings, pongs)
  fmt.Println(<-pongs)
}
