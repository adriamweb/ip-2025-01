package main

import "fmt"

func main (){

	var ant, termos, atual int
	var fibonacci int = 1

	fmt.Print("Insira a quantidade de termos da fibonacci: ")
	fmt.Scan(&termos)

	if termos < 3{
		fmt.Print("Quantidade inválida.")
		return 
	}

	for i:= 0; i < termos; i++{
		fmt.Println(fibonacci)
		ant = atual
		atual = fibonacci
		fibonacci = ant + atual
	}

}