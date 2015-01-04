package main

import (
  "fmt"
  "net/url"
  "strings"
)

// URLs provide a uniform way to locate resources.

func main() {
  p := fmt.Println
  // We'll parse this example URL, which includes scheme,
  // authentication info, host, port, path, query params,
  // and query fragment
  s := "postgres://user:pass@host.com:5432/path?k=v#f"

  // Parse the url and ensure there are no errors
  u, err := url.Parse(s)
  if err != nil {
    panic(err)
  }

  // Accessing the scheme
  p(u.Scheme)

  // User contains all hte authentication info; call
  // Username and Password on this for individual values
  p(u.User)
  p(u.User.Username())
  pass, _ := u.User.Password()
  p(pass)

  // The Host contains both the hostname and the port, if
  // present. Split the Host manually to extract the port
  p(u.Host)
  h := strings.Split(u.Host, ":")
  p(h[0])
  p(h[1])

  // Here we extract the path and the fragment after the #
  p(u.Path)
  p(u.Fragment)

  // To get the query params in a string of k=v format, use
  // RawQuery. You can also parse query params into a map.
  // The parsed query param maps are from strings to slices of
  // strings, so index into [0] if you only want the first value
  p(u.RawQuery)
  m, _ := url.ParseQuery(u.RawQuery)
  p(m)
  p(m["k"][0])
}
