package main

import (
	"fmt"
)

func main(){

	var N, quadI int

	fmt.Print("Insira um valor inteiro: ")
	fmt.Scan(&N)

	if N <= 5 || N >= 2000{
		fmt.Print("VALOR INVALIDO")
	} else{

		for i := 2; i<= N; i = i+2{
			quadI = i*i
			fmt.Printf("%d^2 = %d\n", i, quadI)
		} 
	}

}