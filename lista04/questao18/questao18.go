// Faça um programa que leia um vetor de elementos inteiros de 10 posições, mas já o leia de maneira ordenada
// crescente. Ao final, imprima o vetor.

package main

import (
	"fmt"
	"sort"
)

func main (){

	var numero int
	vetor := make([]int, 0, 10)

	for i:= 0; i < 10; i++{
		fmt.Scan(&numero)

		vetor = append(vetor, numero)
		sort.Ints(vetor)
	}

	fmt.Print(vetor)
}