package main

import "fmt"

func main (){
	var (
		ant, atual int
		Fibonacci = 1
)

	for i:= 0; i < 50; i++{

		fmt.Print(Fibonacci, " ")
		ant = atual
		atual = Fibonacci
		Fibonacci = ant + atual
	}
}