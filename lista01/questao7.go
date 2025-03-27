package main

import "fmt"

func main(){

	var tempFahrenheit, tempCelsius, qtdPolegadas, qtdMilimetros  float32

	fmt.Print("Insira o valor em fahrenheit: ")
	fmt.Scan(&tempFahrenheit)

	fmt.Print("Insira o valor em polegadas: ")
	fmt.Scan(&qtdPolegadas)

	tempCelsius = (5*tempFahrenheit-160)/9
	qtdMilimetros = 25.4*qtdPolegadas

	fmt.Printf("O VALOR EM CELSIUS = %.2f\n", tempCelsius)
	fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f \n",qtdMilimetros)


}