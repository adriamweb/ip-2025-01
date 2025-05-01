package main

import "fmt"

func main() {

	var (
		array        [10]int
		sumPares     int
		countImpares int
	)

	for i:= 0; i < 10; i++{
		fmt.Scan(&array[i])
	}

	for _, valor := range array {
		if valor%2 == 0 {
			sumPares += valor
			fmt.Print(valor, " ")
		} else {
			countImpares++
		}
	}
	fmt.Println("\nSoma dos números pares:", sumPares)

	for _, valor := range array {
		if valor%2 != 0 {
			fmt.Print(valor, " ")
		}
	}

	fmt.Print("\nA quantidade de numeros impares é: ", countImpares)
}
