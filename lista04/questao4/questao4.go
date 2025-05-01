// Faça um programa que leia um vetor A de 10 posições, contendo números inteiros. Determine e mostre, a
// seguir, quais elementos de A estão repetidos e quantas vezes cada um se repete.

package main

import "fmt"

func verificado(x int, lista []int) bool{
	for _, valor := range lista{
		if x == valor{
			return true
		}
	}

	return false

}

func main (){
	var array [10] int
	var conferidos [] int
	var qtd int = 1

	fmt.Print("Insira os valores separados por espaços: ")

	for i:= 0; i< 10; i++{
		fmt.Scan(&array[i])
	}

	for i, valor := range array{
		for i2, j := range array{
			if i != i2 && valor == j && !verificado(valor, conferidos){
				qtd += 1
			}
		}

		if qtd > 1{
			fmt.Printf("O elemento %d se repete %d vezes\n", valor, qtd)
		}
		qtd = 1
		conferidos = append(conferidos, valor)
	}
}