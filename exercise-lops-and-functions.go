package main

import (
	"fmt"
	"math"
)

const Eps = 1e-9

func Sqrt(x float64) float64 {
	z := 1.0
	m := z
	for {
		z = z - (z*z-x)/(2*z)
		if math.Abs(z-m) < Eps {
			break
		}
		m = z
	}
	return z
}

func main() {
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(2))
}
