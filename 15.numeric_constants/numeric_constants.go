package main

import "fmt"

const (
  Big = 1 << 100
  Small = Big >> 99
  Middle = 1 << 1
)

func needInt(x int) int {
  return x*10 + 1
}

func needFloat(x float64) float64 {
  return x*0.1
}

func main() {
  p := fmt.Println
  p(needInt(Small))
  p(needFloat(Small))
  p(needFloat(Big))
  p(needInt(Middle))
}

