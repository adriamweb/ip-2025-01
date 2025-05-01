// Faça um programa que receba 100 valores numéricos, armazene-os em um vetor, calcule e imprima o valor do
// somatório dado a seguir:
// S=(b0-b99)
// 3
// + (b1-b98)
// 3
// + (b2-b97)
// 3
// + . . . + (b49-b50)
// 3

package main

import (
	"fmt"
	"math"
)

func main (){

	var (
		valores[100] float64
		resultados[50] float64
	)

	for i:= 0; i < 100; i++{
		fmt.Scan(&valores[i])
	}

	for i:= 0; i < len(resultados); i ++{
		resultados[i] = math.Pow((valores[i] + valores[len(valores)-1-i]), 3.0)
	}

	fmt.Print(resultados)
}