// Faça um programa que leia um número inteiro positivo na base 10, calcule e imprima o seu equivalente na
// base 2.

package main

import "fmt"

func main() {
	var numero uint64
	var lista []uint64

	fmt.Print("Insira o numero: ")
	fmt.Scan(&numero)

	for numero > 0 {
		lista = append([]uint64{numero%16}, lista...)
		numero /= 16
	}

	for _, valor := range lista {
		fmt.Print(valor)
	}

}
