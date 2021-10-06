package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(p)

	q := p[1:4]
	fmt.Println(q)

	r := p[:2]
	fmt.Println(r)

	s := p[1:]
	fmt.Println(s)
}
