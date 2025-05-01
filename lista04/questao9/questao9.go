package main

import "fmt"

func main() {

	var alturas [10]float64
	var media float64

	for i:= 0; i < 10; i++{
		fmt.Scan(&alturas[i])
	}

	for _, valor := range alturas {
		media += valor
	}

	media /= 10

	for i := 0; i < len(alturas); i++ {
		if alturas[i] > media {
			fmt.Print(alturas[i], " ")
		}
	}
}
