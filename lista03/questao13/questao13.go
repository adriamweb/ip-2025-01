package main

import "fmt"

func main (){

	var contador float64 = 1
	var H float64

	for i:= 1.0; i <= 50; i++{

		H += contador/i
		contador += 2
	}

	fmt.Print("O valor de H é: ", H)
	
}

