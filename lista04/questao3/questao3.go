package main

import "fmt"

func main() {

	var (
		array        [10]int
		sumPares     int
		countImpares int
	)

	fmt.Scan(&array[0], &array[1], &array[2], &array[3], &array[4], &array[5], &array[6], &array[7], &array[8], &array[9])

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
