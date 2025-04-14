package main

import "fmt"

func main() {
	var aiProx, aiAtual, aiAnt, qtdTermos int

	fmt.Print("Insira os dois primeiros numeros da sequencia: ")
	fmt.Scan(&aiAnt, &aiAtual)

	fmt.Print("Insira a quantidade de termos da sequencia que deseja consultar: ")
	fmt.Scan(&qtdTermos)

	fmt.Printf("%d\n%d\n", aiAnt, aiAtual)

	for i := 1; i <= qtdTermos-2; i++ {

		if i%2 != 0 {
			aiProx = aiAtual + aiAnt
			aiAnt = aiAtual
			aiAtual = aiProx
		} else {
			aiProx = aiAtual - aiAnt
			aiAnt = aiAtual
			aiAtual = aiProx
		}

		fmt.Println(aiProx)

	}
}

// A série de Fetuccine é gerada da seguinte forma: os dois primeiros termos (inteiros) são fornecidos pelo
// usuário. A partir daí, os termos são gerados com a soma ou subtração dos dois termos anteriores, ou seja:
// Ai = Ai-1 + Ai-2, para i ímpar;
// Ai = Ai-1 – Ai-2, para i par.
// Crie um programa que imprima os N primeiros termos da série de Fetuccine, assumindo que o usuário digitará
// um N>=3.
