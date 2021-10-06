package main

import "fmt"

func main() {
	khalifas := [4]string{
		"Abu Bakr (R)",
		"Umar (R)",
		"Uthman Bin Affan(R)",
		"Ali Bin Abi Talib(R)",
	}

	a := khalifas[0:2]
	b := khalifas[1:3]

	fmt.Println(a, b)

	b[0] = "Uman bin Khattab (r)"
	fmt.Println(a, b)
	fmt.Println(khalifas)
}
