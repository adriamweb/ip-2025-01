package main

import "fmt"

func main (){
	var numero int

	fmt.Print("Insira um numero inteiro: ")
	fmt.Scan(&numero)

	if numero % 5 == 0{
		fmt.Print("O numero é divisivel por 5.")
	} else{
		fmt.Print("O numero não é divisivel por 5.")
	}
}