package main

import (
  "io/ioutil"
  "log"
  "os"
  // "text/template" also available
  "html/template"
)

type User struct {
  FirstName string
  LastName  string
  Email     string
  Age       int
}

func (u User) IsOld() bool {
  return u.Age > 30
}

func main() {
  u := User{"Daniel", "Carter", "dan@example.com", 26}

  body, _ := ioutil.ReadFile("templates/file1.tmpl")
  htmlbody, _ := ioutil.ReadFile("templates/file2.html")

  tmpl1, err := template.New("Some Name").Parse(string(body))
  if err != nil {
    log.Panic(err)
  }
  tmpl2, err := template.New("Some Name").Parse(string(htmlbody))
  if err != nil {
    log.Panic(err)
  }

  tmpl1.Execute(os.Stdout, u)
  tmpl2.Execute(os.Stdout, u)
}
