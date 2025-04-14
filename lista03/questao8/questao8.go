// Faça um programa que imprima todos os números inteiros de 1 (inclusive) até 20 (inclusive) e também a soma
// de todos eles. 

package main

import "fmt"

func main (){
	var soma int

	for i:= 1; i <= 20; i++{
		fmt.Println(i)
		soma += i
	}

	fmt.Println("A soma é: ", soma)
}