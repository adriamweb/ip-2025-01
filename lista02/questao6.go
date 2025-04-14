package main

import "fmt"

func main() {
	var A, B int

	fmt.Print("Insira o numero inteiro A: ")
	fmt.Scan(&A)

	fmt.Print("Insira o numero inteiro B: ")
	fmt.Scan(&B)

	if A%B == 0 {
		fmt.Printf("O numero %d é divisivel por %d", A, B)
	} else {
		fmt.Printf("O numero %d não é divisivel por %d", A, B)
	}
}
