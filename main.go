package main

import (
  "fmt"
  "os"
)

// use os package to get the env variable which is already set
func envVariable(key string) string {

  // set env variable using os package
  os.Setenv(key, "gopher")

  // return the env variable using os package
  return os.Getenv(key)
}

func main() {
  // os package
  value := envVariable("name")

  fmt.Printf("os package: name = %s \n", value)
  fmt.Printf("environment = %s \n", os.Getenv("APP_ENV"))
}