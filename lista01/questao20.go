package main

import "fmt"

func main () {

	var horas, minutos, segundos, segundosTotais int

	fmt.Print("Insira a quantidade de horas: ")
	fmt.Scan(&horas)

	fmt.Print("Insira a quantidade de minutos: ")
	fmt.Scan(&minutos)

	fmt.Print("Insira a quantidade de segundos: ")
	fmt.Scan(&segundos)
	
	segundosTotais = horas * 60 * 60 + minutos * 60 + segundos

	fmt.Printf("O TEMPO EM SEGUNDOS E = %d\n",segundosTotais)	
}