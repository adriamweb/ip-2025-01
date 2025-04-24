package main

import "fmt"

func main (){
	var num bool = true
	var array [10] int

	fmt.Print("Insira os valores separados por espaços: ")
	fmt.Scan(&array[0], &array[1],&array[2],&array[3],&array[4],&array[5],&array[6],&array[7],&array[8],&array[9])

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