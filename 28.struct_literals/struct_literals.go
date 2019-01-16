package main

import "fmt"

type Vertex struct {
  X, Y int
}

var (
  p = Vertex{1,2}
  q = &Vertex{1,2}
  r = Vertex{X:1}  // Y:0 implict
  s = Vertex{}  // X:0, Y:0 implict
)

func main() {
  fmt.Printf("%v\n%v\n%v\n%v\n\n", p, q, r, s)
  fmt.Printf("%+v\n%+v\n%+v\n%+v\n", p, q, r, s)
}

