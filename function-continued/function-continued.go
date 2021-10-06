package main

import (
	"fmt"
)

func addShortened(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(addShortened(100, 200))
}
