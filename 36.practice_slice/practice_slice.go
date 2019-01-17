package main

import (
  "golang.org/x/tour/pic"
)

func Pic (dx, dy int) [][]uint8 {
  a := make([][]uint8, dy)
  for i := range a {
    a[i] = make([]uint8, dx)
    for x := range a[i] {
      a[i][x] = uint8((x+i)/2)
    }
  }
  return a
}

func main() {
  pic.Show(Pic)
}
