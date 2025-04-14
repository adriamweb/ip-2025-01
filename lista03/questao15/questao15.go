package main

import "fmt"

func main() {

	var (
		nSequencia    int
		valor         int = 1
		incrementador int = 3
	)

	fmt.Print("Insira a quantidade de termos da sequencia: ")
	fmt.Scan(&nSequencia)

	for i := 0; i < nSequencia; i++ {
		fmt.Print(valor, " ")
		valor += incrementador
		incrementador += 2
	}

}