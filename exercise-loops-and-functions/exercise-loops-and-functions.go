package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := 1.0
	for math.Abs(z*z-x) > 0.000001 {
		z = z - (z*z-x)/(2*z)
		fmt.Println(z)
	}

	return z
}

func main() {
	fmt.Println(sqrt(6))
	fmt.Printf("The distance from standard lib sqrt is %v", math.Sqrt(6)-sqrt(6))
}
