package main

import (
	"fmt"
	"math"
)


func main (){
	raiz := [15] float64 {}
	valores := [15] int {}

	for i:= 0; i < 15; i++{
		fmt.Scan(&valores[i])
	}

	for i, valor := range valores{
		if valor < 0{
			raiz[i] = -1
		} else{
			raiz[i] = math.Sqrt(float64(valor))
		}
	}


	for _, valor := range raiz{
		fmt.Printf("%.2f ", valor)
	}

}