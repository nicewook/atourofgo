package main

import (
  "fmt"
)

func swap(x, y string) (string, string) {  // can have multi returns
  return y, x
}

func main() {
  a, b := swap("AAAAA back", "BBBBB front")
  fmt.Println(a, b)
}

