package main

import "fmt"

func main(){

	var nota float32

	fmt.Print("Insira a nota: ")
	fmt.Scan(&nota)


	fmt.Printf("NOTA: %.1f ", nota)
	if  9 <= nota && nota <= 10{
		fmt.Printf("CONCEITO = A\n")
	} else if 7.5 <= nota && nota < 9{
		fmt.Printf("CONCEITO = B\n")
	} else if 6 <= nota && nota < 7.5{
		fmt.Printf("CONCEITO = C\n")
	} else if 0 <= nota && nota < 6 {
		fmt.Print("CONCEITO = D\n")
	}

}