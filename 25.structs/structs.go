package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  fmt.Println(Vertex{1,2})
  my_vertex := Vertex{X:3, Y:4}
  fmt.Printf("%v\n%#v\n%+v\n",my_vertex, my_vertex, my_vertex)
}


