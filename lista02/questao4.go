package main

import ("fmt"
		"math")

func main(){
	var numero float64

	fmt.Print("Insira um numero float: ")
	fmt.Scan(&numero)

	if numero < 0{
		fmt.Print(numero * numero)
	} else{
		fmt.Print(math.Pow(numero, 0.5))
	}
}