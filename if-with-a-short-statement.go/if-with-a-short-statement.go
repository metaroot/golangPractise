package main

import (
	"fmt"
	"math"
)

func limitedPow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(limitedPow(3, 2, 10))
	fmt.Println(limitedPow(3, 3, 10))

}
