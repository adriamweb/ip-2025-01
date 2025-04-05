package main

import "fmt"

func main (){

	var numero int

	fmt.Print("Insira um numero inteiro: ")
	fmt.Scan(&numero)

	if numero > 20 && numero < 90{
		fmt.Print("O numero está entre 20 e 90.")
	} else{
		fmt.Print("O numero não está entre 20 e 90.")
	}
}