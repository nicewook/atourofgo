package main

import (
  "fmt"
  "math/cmplx"
)

var (
  ToBe bool = false
  MaxInt uint64 = 1<<64 - 1
  z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
  const f = "%T(%v)\n"
  p := fmt.Printf
  p(f, ToBe, ToBe)
  p(f, MaxInt, MaxInt)
  p(f, z, z)
}

