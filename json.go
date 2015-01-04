package main

import (
  "encoding/json"
  "fmt"
  "os"
)

// Go offers built-in support for JSON encoding and
// decoding, including to and from built-in and custom
// data types.

// We'll use these two structs to demonstrate encoding
// and decoding of custom types
type Response1 struct {
  Page int
  Fruits []string
}
type Response2 struct {
  Page   int      `json:"page"`
  Fruits []string `json:"fruits"`
}

func main() {
  p := fmt.Println
  //First, we'll look at encoding basic data types to
  // JSON strings. Here are some examples for atomic values.
  bolB, _ := json.Marshal(true)
  p(string(bolB))

  intB, _ := json.Marshal(1)
  p(string(intB))

  fltB, _ := json.Marshal(2.34)
  p(string(fltB))

  strB, _ := json.Marshal("gopher")
  p(string(strB))

  // Here are some examples for slices and maps, which
  // encode to JSON arrays and objects as you'd expect
  slcD := []string{"apple", "peach", "pear"}
  slcB, _ := json.Marshal(slcD)
  p(string(slcB))

  mapD := map[string]int{"apple": 5, "lettuce": 7}
  mapB, _ := json.Marshal(mapD)
  p(string(mapB))

  // The JSON package can automatically encode your custom
  // data types. It will only include exported fields in the
  // encoded outout and will by default use those names as the
  // JSON keys.
  res1D := &Response1{
    Page: 1,
    Fruits: []string{"apple", "peach", "pear"}}
  res1B, _ := json.Marshal(res1D)
  p(string(res1B))

  // You can use tags on struct field declarations to customize
  // the encoded JSON key names. Check the definition of
  // Response2 above to see an example of such tags.
  res2D := &Response2{
    Page: 1,
    Fruits: []string{"apple", "peach", "pear"}}
  res2B, _ := json.Marshal(res2D)
  p(string(res2B))

  // Now let's look at decoding JSON data into Go values.
  byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

  // We need to provide a variable where the JSON package
  // can put the decoded data. This map[string]interface{}
  // will hold a map of strings to arbitrary data types.
  var dat map[string]interface{}

  // Here's the actual decoding, and a check for associated
  // errors
  if err:= json.Unmarshal(byt, &dat); err != nil {
    panic(err)
  }
  p(dat)

  // In order to use the values in the decoded map, we'll need
  // to cast them to their appropriate types. For example, here
  // we cast the value in num to the expected float64 type
  num := dat["num"].(float64)
  p(num)

  // Accessing nested data requires a series of casts
  strs := dat["strs"].([]interface{})
  str1 := strs[0].(string)
  p(str1)

  // We can also decode JSON into custom data types. This has
  // the advantage of adding additional type-safety to out
  // programs and eliminating the need for type assertions
  // when accessing the decoded data.
  str := `{"page": 1, "fruits": ["apples", "peach"]}`
  res := &Response2{}
  json.Unmarshal([]byte(str), &res)
  p(res)
  p(res.Fruits[0])

  // In the example above, we always used bytes and strings as
  // intermediates between the data and JSON representation on
  // strout. We can also stream JSON encodings directly to
  // os.Writers like os.Stdout or even HTTP response bodies
  enc := json.NewEncoder(os.Stdout)
  d := map[string]int{"apple": 5, "lettuce": 7}
  enc.Encode(d)
}
