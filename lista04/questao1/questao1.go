package main

import "fmt"

func main (){
	var num bool = true
	var array [10] int

	fmt.Print("Insira os valores separados por espaços: ")

	for i:= 0; i < 10; i++{
		fmt.Scan(&array[i])
	}

	for indice, valor := range array{
		if valor > 50{
			fmt.Printf("Valor: %d Indice: %d\n", valor, indice)
			num = false
		}

	}

	if num{
		fmt.Print("Nao existe valor superior a 50.")
	}
}