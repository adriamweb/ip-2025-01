package main

import "fmt"

func main(){

	var qtdConversoes int
	var tempFahrenheit, tempCelsius float32

	fmt.Print("Insira a quantidade de temperaturas fahrenheit a serem convertidas: ")
	fmt.Scan(&qtdConversoes)
	var valores [][]float32


	for i:= 1; i <= qtdConversoes; i = i+1{
		fmt.Printf("Insira a %d° temperatura em fahrenheit: ", i)
		fmt.Scan(&tempFahrenheit)
		tempCelsius = 5*(tempFahrenheit-32)/9

		valores = append(valores, []float32{tempCelsius, tempFahrenheit})

	}

	for _, valor:= range valores{
		fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", valor[1],valor[0])

	}

}