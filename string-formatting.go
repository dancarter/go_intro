package main

import "fmt"
import "os"

// Go offers excellent support for string formattin in the
// printf tradition. Here are some examples of common string
// formatting tasks.

type point struct {
  x, y int
}

func main() {
  p := point{1, 2}

  // Go offers several printing 'verbs' designed to format
  // general Go values. For example, this prints an instance
  // of our point struct
  fmt.Printf("%v\n", p)

  // If the value is a struct, the '%+v' variant will include
  // the struct's field names
  fmt.Printf("%+v\n", p)

  // The '%#v' variant prints a Go syntax representation of the
  // value, ie the source code snippet that would produce that
  // value.
  fmt.Printf("%#v\n", p)

  // The type of a value is printed with '%T'
  fmt.Printf("%T\n", p)

  // '%t' for booleans
  fmt.Printf("%t\n", true)

  // There are many options for integers. Use '%d' for standard,
  // base-10 formatting
  fmt.Printf("%d\n", 123)

  // Use '%b' fir a binary representation
  fmt.Printf("%b\n", 123)

  // Use '%c' to print the character corresponding to the given
  // integer.
  fmt.Printf("%c\n", 33)

  // Use '%x' for hex encoding
  fmt.Printf("%x\n", 456)

  // Floats also have several formatting options. Use '%f' for
  // basic decimal formatting
  fmt.Printf("%f\n", 78.9)

  // '%e' and '%E' format the float in (slightlt different versions
  // of) scientific notation.
  fmt.Printf("%e\n", 123400000.0)
  fmt.Printf("%E\n", 123400000.0)

  // Use '%s' for basic string printing
  fmt.Printf("%s\n", "\"string\"")

  // To double quote strings, use '%q'
  fmt.Printf("%q\n", "string")

  // As with integers, '%x' renders the string in base-16, with two
  // output characters per byte of input
  fmt.Printf("%x\n", "hex this")

  // '%p' prints a representation of a point
  fmt.Printf("%p\n", &p)

  // When formatting numbers, you will often want to control the
  // width and percision of the resulting figure. To specify the
  // width of an integer, use a number after the % in the verb. By
  // default, the result will be right-justified and padded with
  // spaces.
  fmt.Printf("|%6d|%6d|\n", 12, 345)

  // You can also specify the width of printed floast, though usually
  // you'll also want to restrict the decimal precision at the same
  // time with the 'width.precision' syntax
  fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

  // To left-justify, use the '-' flag
  fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

  // You may also want to control width when formatting strings,
  // especially to ensure that they align in table-like output.
  // For basic right-justified width
  fmt.Printf("|%6s|%6s|\n", "foo", "b")

  // To left-justify, use the '-' flag
  fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

  // 'Sprintf' formats and returns a string without printing it
  s := fmt.Sprintf("a %s", "string")
  fmt.Println(s)

  // You can format+print to io.Writers other than os.Stdout
  // using Fprintf
  fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
