package main

import (
	"fmt"
	"math"
)

// Abser exported
type Abser interface { // export
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v
	//a = v

	fmt.Println(a.Abs())

}

// MyFloat exported type
type MyFloat float64

// Abs is...
func (f MyFloat) Abs() float64 { // export
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Vertex exported type
type Vertex struct { // export
	X, Y float64
}

// Abs exported fucntion
func (v *Vertex) Abs() float64 { // export
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
