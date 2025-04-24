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
	fmt.Scan(&vetor1[0], &vetor1[1],&vetor1[2],&vetor1[3],&vetor1[4],&vetor1[5],&vetor1[6],&vetor1[7],&vetor1[8],&vetor1[9])

	fmt.Print("Insira os valores do vetor 2: ")
	fmt.Scan(&vetor2[0], &vetor2[1],&vetor2[2],&vetor2[3],&vetor2[4])
	
	for _, valor := range vetor1{

		if valor %2 == 0{
			sumPar += valor
		} else{
			sumImpar += valor
		}
	}

	fmt.Print(sumPar, sumImpar)
	for i, _ := range vetor2{
		resultante1[i] = vetor2[i] + sumPar
		resultante2[i] = vetor2[i] + sumImpar
	}
	
	fmt.Println(resultante1)
	fmt.Println(resultante2)
}