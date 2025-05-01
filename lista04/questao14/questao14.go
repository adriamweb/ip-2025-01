package main

import "fmt"

func main() {
	var (
		vetor1, vetor2 [10]int
		intercalado    []int
	)

	fmt.Print("Insira os valores do 1° vetor: ")

	for i := 0; i < 10; i++ {
		fmt.Scan(&vetor1[i])
	}

	fmt.Print("Insira os valores do 2° vetor: ")

	for i := 0; i < 10; i++ {
		fmt.Scan(&vetor2[i])
	}

	for i := 0; i < 10; i++ {
		intercalado = append(intercalado, vetor1[i], vetor2[i])
	}

	fmt.Print("Vetor intercalado: ", intercalado)

}
