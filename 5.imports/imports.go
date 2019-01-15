package main 

import (
  "fmt"
  "math"
)

func main() {
  p := fmt.Printf
  p("Now you have %g problems.", math.Nextafter(2,3))
  // Nextafter returns the next representable float64 value after x towards y.
}
