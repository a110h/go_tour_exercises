package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	
	minDiff := float64(0.0001)
	z := float64(1)
	z_last := z
	for i := 0; i < 100; i++ {
		z -= (z*z - x) / (2 * z)		
		
		diff := float64(z-z_last)
		fmt.Println(z, z_last, diff)	

		if (diff < 0 && diff >-minDiff) || (diff > 0 && diff < minDiff) {
			break
		} else {
			z_last = z
		}
	}

	return z, nil
}


func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
