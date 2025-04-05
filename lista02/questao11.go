package main

import "fmt"

func main (){
	var x float64

	fmt.Print("Insira o valor de X: ")
	fmt.Scan(&x)

	fmt.Printf("O valor de f(x) é: %.2f", (8 /(2-x)))
}