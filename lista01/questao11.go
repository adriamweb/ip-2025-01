package main

import "fmt"

func main(){

	var numero int

	fmt.Print("Insira um numero inteiro: ")
	fmt.Scan(&numero)

	if numero % 3 == 0 && numero % 5 == 0{
		fmt.Print("O NUMERO E DIVISIVEL\n")
	} else {
		fmt.Print("O NUMERO NAO E DIVISIVEL\n")
	}

}