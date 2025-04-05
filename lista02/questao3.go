package main

import "fmt"

func main (){
	var num1, num2 int

	fmt.Print("Insira o primeiro numero: ")
	fmt.Scan(&num1)

	fmt.Print("Insira o segundo numero: ")
	fmt.Scan(&num2)

	if num1 + num2 > 20{
		fmt.Print(num1 + num2 + 8)
	} else{
		fmt.Print(num1 + num2 - 5)
	}
}