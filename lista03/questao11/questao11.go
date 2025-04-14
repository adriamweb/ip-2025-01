package main

import "fmt"

func main (){
	var qtdFatorial int
	var fatorial int = 1

	fmt.Print("Insira o numero que deseja o fatorial: ")
	fmt.Scan(&qtdFatorial)

	if fatorial < 0{
		fmt.Println("Valor inválido!")
		return
	}

	for i:= qtdFatorial; i > 0; i = i - 1{
		fatorial *= i
	}

	fmt.Print("O fatorial é: ", fatorial)
}