package main

import "fmt"

func main (){
	var numero, dezena int

	fmt.Print("Insira um número que possua 3 casas: ")
	fmt.Scan(&numero)

	if numero < 100 || numero > 999{
		fmt.Print("Numero inválido.")
		return
	}

	dezena = (numero/10)%10
	fmt.Print("A unidade de dezena é: ", dezena)
}