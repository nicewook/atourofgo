package main

import "fmt"

type Vertex struct {
  X, Y int
}

func main() {
  v := new(Vertex)  // v is pointer
  fmt.Printf("%v\n%+v\n%#v\n", v, v, v)

  v.X, v.Y = 11, 9
  fmt.Printf("%v\n%+v\n%#v\n", v, v, v)
}
