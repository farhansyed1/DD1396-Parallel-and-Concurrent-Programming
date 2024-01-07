package main

import (
	"fmt"
	"math"
)

// Sqrt find an approximate value of the square root of a number
func Sqrt(x float64) float64 {
	z := 1.0
	for {
		previousZ := z
		z -= (z*z - x) / (2 * z)
		if math.Abs(previousZ-z) < 0.00001 {
			return z
		}
	}
}

func main() {
	newtonsMethod := Sqrt(2)
	realValue := math.Sqrt(2)
	fmt.Printf("Newton's method: %g \n", newtonsMethod)
	fmt.Printf("Standard library function: %g \n", realValue)
	fmt.Printf("Difference: %g", math.Abs(newtonsMethod-realValue))
}
