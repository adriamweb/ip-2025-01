package main

import "fmt"

func main (){
	var numero int

	fmt.Print("Insira um numero inteiro: ")
	fmt.Scan(&numero)

	if numero % 2 == 0{
		fmt.Print("O numero é par.")
	} else{
		fmt.Print("O numero é impar.")
	}
}