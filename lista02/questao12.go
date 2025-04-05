package main

import "fmt"

func main (){
	var idade int

	fmt.Print("Insira uma idade: ")
	fmt.Scan(&idade)

	if idade < 1{
		fmt.Print("Idade inválida. ")
		return
	}

	if 0 <= idade && idade <= 2{
		fmt.Print("Recém-nascido")
	} else if 3 <= idade && idade <= 11{
		fmt.Print("Criança")
	} else if 12 <= idade && idade <= 19{
		fmt.Print("Adolescente")
	} else if 20 <= idade && idade <= 55{
		fmt.Print("Adulto")
	} else{
		fmt.Print("Idoso")
	}
}