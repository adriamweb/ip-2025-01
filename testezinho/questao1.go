package main

import "fmt"

func main (){
	var num int

	fmt.Scan(&num)

	if num > 0{
		fmt.Print("Numero positivo")
	} else if num < 0{
		fmt.Print("Numero negativo")
	} else{
		fmt.Print("Nulo")
	}
}