package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Printf("\n%d° coluna: ", i)
		for j := 0; j < 10; j++ {
			fmt.Print(j, " ")
		}
	}
}

// . Escreva um programa que imprima os índices de todos os elementos de uma matriz 10x10.
