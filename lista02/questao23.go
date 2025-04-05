package main

import "fmt"

func main (){
	var idade int

	fmt.Print("Insira a idade: ")
	fmt.Scan(&idade)

	if idade < 16{
		fmt.Print("Não eleitor")
	} else if idade >= 18 && idade <= 65{
		fmt.Print("Eleitor obrigatório")
	} else{
		fmt.Print("Eleitor facultativo")
	}
}