package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(f float64) float64 {
	return f * 0.01
}

func main() {
	fmt.Println(needFloat(Big))
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
}
