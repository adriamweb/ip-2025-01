package main

import "fmt"

func main (){

	var x, y int

	fmt.Print("Insira dois numeros inteiros: ")
	fmt.Scan(&x,&y)

	if x%2 == 0{
		y = y*2 + x
		for i:= x; i < y; i = i+2{
			fmt.Print(i," ")
		}
		fmt.Print("\n")
	} else {
		fmt.Print("O PRIMEIRO NUMERO NAO E PAR")
	}
}