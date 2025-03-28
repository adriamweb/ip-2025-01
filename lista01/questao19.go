package main

import "fmt"

func main (){

	var (
		numero, somatorio float64
	)

	fmt.Print("Insira um numero inteiro positivo: ")
	fmt.Scan(&numero)
	numeroVerify := int(numero)

	if numero < 1 || float64(numeroVerify) != numero {
		fmt.Print("Numero invalido!")
	} else {
		
		for i:= 1.0; i <= float64(numeroVerify); i = i+1{
			somatorio = somatorio + 1.0/i
		}
		fmt.Printf("%.6f",somatorio)
		
	}

}