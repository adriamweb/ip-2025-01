package main

import "fmt"

func main (){
	var N, somatorio int

	fmt.Print("Insira o numero: ")
	fmt.Scan(&N)

	for i:=1; i <= N; i++{
		somatorio += i
	}

	fmt.Print("\nO somatório é: ", somatorio)
}