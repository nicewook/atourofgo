package main

import (
  "fmt"
)

func main() {
  q := []int{2,3,5,7,11,13}
  p := fmt.Println

  p("q ==", q)
  p("q[1:4] ==", q[1:4])
  
  p("q[:3] ==", q[:3])
  p("q[4:] ==", q[4:])
}


