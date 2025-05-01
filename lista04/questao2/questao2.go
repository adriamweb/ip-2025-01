package main

import (
	"fmt"
)

func main (){
	var (
		vetor1 [10] int
		vetor2 [5] int
		resultante1[5] int
		resultante2[5] int
		sumPar int
		sumImpar int
	)

	fmt.Print("Insira os valores do vetor 1: ")
	for i:= 0; i < 10; i++{
		fmt.Scan(&vetor1[i])
	}

	fmt.Print("Insira os valores do vetor 2: ")
	for i := 0; i < 5; i++{
		fmt.Scan(&vetor2[i])
	}
	
	for _, valor := range vetor1{

		if valor %2 == 0{
			sumPar += valor
		} else{
			sumImpar += valor
		}
	}

	for i, _ := range vetor2{
		resultante1[i] = vetor2[i] + sumPar
		resultante2[i] = vetor2[i] + sumImpar
	}
	
	fmt.Println(resultante1)
	fmt.Println(resultante2)
}