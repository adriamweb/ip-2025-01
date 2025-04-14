// Escreva um programa que receba um número inteiro positivo, verifique e informe se ele é ou não um número
// triangular. Obs.: Um número é triangular quando é resultado do produto de três números naturais consecutivos.
// Exemplo: 24 = 2 x 3 x 4; 120 = 4 x 5 x 6

package main

import "fmt"


func main (){
	var (
		numero int
		contador1 int
		contador2 int = 1
		contador3 int = 2
	)

	fmt.Print("Insira um numero inteiro: ")
	fmt.Scan(&numero)

	for contador1 * contador2 * contador3 < numero{
		contador1 ++
		contador2 ++
		contador3 ++
	}

	if contador1 * contador2 * contador3 == numero{
		fmt.Print("O numero é triangular!")
	} else{
		fmt.Print("O numero não é triangular!")
	}
}