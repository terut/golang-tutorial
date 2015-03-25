package main

import (
	"fmt"
	"math"
)

const Eps = 1e-9

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	m := z
	for {
		z = z - (z*z-x)/(2*z)
		if math.Abs(z-m) < Eps {
			break
		}
		m = z
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", e)
}
