package main

import "fmt"


// Escreva um programa que armazene os 100 primeiros números ímpares em um vetor. Imprima o vetor em
// seguida.

func main (){
	vetorImpar := [] int {}

	for i:=1; len(vetorImpar) < 100; i+= 2{
		vetorImpar = append(vetorImpar, i)
	}

	fmt.Print(vetorImpar)
}