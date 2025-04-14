package main

import "fmt"

func main() {
	var (
		soma     float64
		contador float64 = 1
	)

	for i := 38.0; i-1 >= 1; i = i - 1 {
		soma += (i * (i - 1) / contador)
		contador++
	}

	fmt.Print("A soma é: ", soma)
}
