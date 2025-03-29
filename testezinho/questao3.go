package main

import "fmt"

func main (){
	var soma int
	var contador int = 1

	for contador <= 100{
		fmt.Print(contador," ")
		soma += contador
		contador += 1
	}

	fmt.Print("\nA soma é: ", soma)
}