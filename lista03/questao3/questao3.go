package main

import "fmt"

func main (){
	var (salarioCarlos float64
		contador int)

	fmt.Print("Insira o salário de Carlos: ")
	fmt.Scan(&salarioCarlos)

	salarioJoao:= salarioCarlos/3

	for salarioJoao < salarioCarlos{
		salarioCarlos*= 1.02
		salarioJoao*= 1.05
		contador += 1
		}

	fmt.Printf("A quantidade de meses necessários é: %d", contador)
}