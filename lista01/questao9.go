package main 

import (
	"fmt"
)


func main(){

	var A, B, C, delta float32

	fmt.Print("Insira o valor do coeficiente A: ")
	fmt.Scan(&A)

	fmt.Print("Insira o valor do coeficiente B: ")
	fmt.Scan(&B)

	fmt.Print("Insira o valor do coeficiente C: ")
	fmt.Scan(&C)

	delta = B * B - 4*A*C

	fmt.Printf("O VALOR DE DETAL E = %.2f", delta)
	
}