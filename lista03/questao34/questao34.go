package main

import "fmt"

func primo(numero int) bool {

	for i := 2; i*i <= numero; i++ {
		if numero%i == 0 {
			return false
		}
	}

	return true
}

func proximoPrimo(numero int) int {
	for {
		numero++
		if primo(numero) {
			return numero
		}
	}
}

func main() {

	var N1, N2 int
	var primo, multiplo int = 2, 1

	fmt.Print("Insira os numeros: ")
	fmt.Scan(&N1, &N2)

	for N1 != 1 || N2 != 1 {

		for (N1%primo == 0 && N1/primo > 0) || (N2%primo == 0  && N2/primo > 0){

			if N1%primo == 0{
				N1 /= primo
			}

			if  N2 % primo == 0{
				N2 /= primo
			}

			multiplo *= primo
		}
		primo = proximoPrimo(primo)
	}

	fmt.Print("\nO menor multiplo comum é: ", multiplo)

}
