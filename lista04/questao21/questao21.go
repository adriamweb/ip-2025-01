package main

import "fmt"

func main() {
	var (
		vetor  [10]int
		codigo int
	)

	for i := 0; i < len(vetor); i++ {
		fmt.Scan(&vetor[i])
	}

	fmt.Print("Insira o código: ")
	fmt.Scan(&codigo)

	switch codigo {

	case 0:
		return
	case 1:
		fmt.Print(vetor)
	case 2:
		for i := len(vetor)-1; i >= 0; i-- {
			fmt.Print(vetor[i], " ")
		}
	default:
		fmt.Print("Código inválido")

	}
}
