package main

import "fmt"

func main(){

	var salario float32

	fmt.Print("Insira o salario: ")
	fmt.Scan(&salario)

	if salario <= 300{
		salario = salario *1.5
	} else{
		salario = salario * 1.3
	}

	fmt.Printf("SALARIO COM REAJUSTE: %.2f\n",salario)

}