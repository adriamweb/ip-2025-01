package main

import (
	"fmt"
	"math"
)

func fatorial(n float64) float64 {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}

func cosseno(x float64) float64 {
	var soma float64 = 1
	var expoente float64 = 2

	for i := 1; i <= 20; i++ {
		if i%2 != 0 {
			soma -= (math.Pow(x, expoente) / fatorial(expoente))
		} else {
			soma += (math.Pow(x, expoente) / fatorial(expoente))
		}

		expoente += 2
	}

	return soma
}

func main() {
	var X float64

	fmt.Print("Insira o X do para calcular o cosseno: ")
	fmt.Scan(&X)

	cossenoAprox := cosseno(X)
	cossenoExato:= math.Cos(X)

	fmt.Printf("A diferença entre os cossenos é: %.5f\n", cossenoAprox- cossenoExato)
	fmt.Printf("\nItem A: %.5f\nItem B: %.5f", cossenoAprox, (cossenoAprox- cossenoExato))
}
