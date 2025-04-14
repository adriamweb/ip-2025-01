package main

import "fmt"

func main() {
	var N1, N2, Valor int

	fmt.Print("Insira os dois numeros: ")
	fmt.Scan(&N1, &N2)

	for N2 > 0{
		Valor += N1
		N2-= 1
	}

	fmt.Print(Valor)

}