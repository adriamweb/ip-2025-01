package main

import "fmt"

func main() {
	var (
		fatorial float64 = 1
		numero   float64 = 100
		soma     float64
	)

	for i := 1.0; i < 20; i++ {
		fatorial *= i
		soma += (numero - i) / fatorial
	}

	fmt.Print("A soma é: ", soma)
}
