package main

import (
	"fmt"
	"math/cmplx"
)

func cbrt(x complex128) complex128 {
	z := complex128(2)
	s := complex128(0)
	for {
		z = z - (cmplx.Pow(z, 3)-x)/(3*z*z)
		if cmplx.Abs(s-z) < 1e-7 {
			break
		}
		s = z
	}
	return z

}

func main() {
	fmt.Println(cbrt(2))
}
