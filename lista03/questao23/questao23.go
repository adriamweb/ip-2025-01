package main

import "fmt"

func main() {

	var (
		qtdTermos int
		contador  float64 = 1
		soma float64 = 0
	)

	fmt.Print("Insira a quantidade de termos desejada: ")
	fmt.Scan(&qtdTermos)

	for i := -1000.0; i < (-1000.0 + float64(qtdTermos * 3)); i = i + 3 {

		if int(contador) %2 != 0{
			soma += -1.0*(i/contador)
		} else{
			soma += (i/contador)
		}

		contador ++

	}

	fmt.Print("\nO valor da soma é: ", soma)
}
