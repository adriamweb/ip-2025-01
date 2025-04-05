package main

import "fmt"

func main (){
	var numero int

	fmt.Print("Insira um numero inteiro: ")
	fmt.Scan(&numero)

	if numero > 0{
		fmt.Print("O numero é positivo.")
	} else if numero < 0{
		fmt.Print("O numero é negativo.")
	} else{
		fmt.Print("O numero é nulo.")
	}
}