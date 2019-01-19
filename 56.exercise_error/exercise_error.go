package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot support complex number: %f.", e)
}

const e = 1e-8

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeSqrt(f)
	}
	return math.Sqrt(f), nil
}

func main() {
	p := fmt.Println
	p(Sqrt(2))
	p(Sqrt(-2))
}
