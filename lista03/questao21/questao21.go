package main

import "fmt"

//b ^ n; n > 1; b>= 2

func Potencia(base int, expoente int) int{
	for i:= 1; i < expoente; i++{
		base*= base
	}
	return base
}

func main (){
	var base, expoente int

	fmt.Print("Insira a base e o expoente, respectivamente: ")
	fmt.Scan(&base, &expoente)

	if base >= 2 && expoente > 1{
		fmt.Print(Potencia(base, expoente))
	} else{
		fmt.Print("Valor inválido")
		return
	}
	
}