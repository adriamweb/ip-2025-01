package main

import "fmt"

//par dobro; impar triplo

func main (){
	
	var vetor, gerado [30] int

	for i:= 0; i < 30; i++{
		fmt.Scan(&vetor[i])
	}

	for i:= 0; i < 30; i++{

		if i% 2 == 0{
			gerado[i] = vetor[i] * 2
		} else{
			gerado[i] = vetor[i] * 3
		}
	}

	fmt.Print(gerado)
}