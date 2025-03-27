package main

import "fmt"

func main (){

	var hora, total int

	fmt.Print("Insira a quantidade de horas em que a charrete foi usada: ")
	fmt.Scan(&hora)

	total = (hora/3)*10 + (hora % 3)*5
	
	fmt.Printf("O VALOR A PAGAR E = %.2d\n",total)
	
}