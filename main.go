package main

import (
  "fmt"
  "github.com/E-Health/gocdm/model"
  "github.com/E-Health/gocdm/api"
)
func main() {
  yob := model.Person{YearOfBirth: 2010}
  name := api.GetNameForTest()
  fmt.Println(yob.TableName(), name)
}