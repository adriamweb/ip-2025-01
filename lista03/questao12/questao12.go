package main

import "fmt"

func main() {

	var X float64
	var fatorial int = 1

	fmt.Print("Insira um valor: ")
	fmt.Scan(&X)
	soma := X


	for i := 1; i <= 20; i = i + 1 {
		fatorial *= i* (-1)
		soma += X/float64(fatorial)
	}

	fmt.Print("A soma é: ", soma)
}
