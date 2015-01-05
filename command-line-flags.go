package main

import "flag"
import "fmt"

// Command-line flags are a common way to specify options
// for command-line programs. For example, in 'wc -l' the
// -l is a command-line flag

// Go provides a flag package supporting basic command-line
// flag parsing. We'll use this package to implement our
// example command-line program.

func main() {
  // Basic flag declarations are available for string,
  // integer, and boolean options. Here we declare a string
  // flag word with a default value "foo" and a short
  // description. This flag.String function returns a string
  // pointer (not a string value)
  wordPtr := flag.String("word", "foo", "a string")

  // This declares numb and fork flags, using a similar
  // approach to the word flag
  numbPtr := flag.Int("numb", 42, "an int")
  boolPtr := flag.Bool("fork", false, "a bool")

  // It's also possible to declare an option that uses an
  // existing var declared elsewhere in the program. Note
  // that we need to pass in a pointer to the flag
  // declaration function.
  var svar string
  flag.StringVar(&svar, "svar", "bar", "a string var")

  // Once all flags are declared, call flag.Parse() to
  // execute the command-line parsing
  flag.Parse()

  // Here we'll just dump out the parsed options and any
  // trailing positional arguments. Note that we need to
  // dereference the points with eg *wordPtr to get the
  // actual option values.
  p := fmt.Println
  p("word:", *wordPtr)
  p("numb:", *numbPtr)
  p("fork:", *boolPtr)
  p("svar:", svar)
  p("tail:", flag.Args())

  // To experiment with the command-line flags program it's
  // best to first compile it and then run the resulting
  // binary directly
  // $ go build command-line-flags.go
  // $ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
  // $ ./command-line-flags -word=opt
  // $ ./command-line-flags -word=opt a1 a2 a3
  // $ ./command-line-flags -h
  // $ ./command-line-flags -wat
}
