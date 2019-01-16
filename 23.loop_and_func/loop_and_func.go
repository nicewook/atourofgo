package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) {
  math_sqrt_x := math.Sqrt(x)
  p := fmt.Printf
  fmt.Println(math_sqrt_x)

  z := float64(1)
  for i := 10; i > 0; i-- {
    z = z - (z * z - x) / (2 * z)
  }
  p("z 10 times: %g\n", z)

  z = float64(1)
  j := 0
  for {
    j++
    z = z - (z * z - x) / (2 * z)
    if math.Abs(z - math_sqrt_x) < 0.0001 {
      break
    }
  }
  p("z %d times: %g\n", j, z)
}

func main() {
  for i:=0; i<10; i++ {
    fmt.Printf("\nsqrt %d\n", i)
    Sqrt(float64(i))
  }
}



