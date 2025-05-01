package main

import "fmt"

func Min(vetor [10]int) (int, int) {

	min := vetor[0]
	var indice int

	for i, valor := range vetor {
		if min > valor {
			min = valor
			indice = i
		}
	}

	return min, indice
}

func main() {

	var array [10]int

	for i:= 0; i < 10; i++{
		fmt.Scan(&array[i])
	}

	menor, indice := Min(array)

	fmt.Printf("O menor elemento do vetor é: %d e sua posição dentro do vetor é: %d", menor, indice)
}
