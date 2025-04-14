package main

import (
	"fmt"
	"math"
)

func main() {
	var numero int

	for numero >= 0 {
		fmt.Print("Insira o número: ")
		fmt.Scan(&numero)

		if float64(int(math.Pow(float64(numero), 0.5))) == float64(math.Pow(float64(numero), 0.5)) {
			fmt.Print("É quadrado perfeito!\n")
		} else {
			fmt.Print("Não é quadrado perfeito!\n")
		}
	}
}
