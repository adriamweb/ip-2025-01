// Elabore um programa que apresente os resultados da soma e da média aritmética dos valores pares situados na
// faixa numérica de 50 a 70.

package main

import "fmt"

func main() {
	var soma, contador float64

	for i := 50.0; i <= 70.0; i += 2 {
		contador += 1
		soma += i
	}

	fmt.Printf("A soma: %.1f\n", soma)
	fmt.Printf("A media: %.1f", soma/contador)
}
