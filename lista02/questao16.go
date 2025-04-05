package main

import ("fmt"
		"math")

func main (){
	var A, B,C float64

	fmt.Print("Insira, separado por espaços, os coeficientes A, B e C, respectivamente: ")
	fmt.Scan(&A, &B, &C)

	delta := math.Pow(B,2) - 4*A*C

	if delta < 0{
		fmt.Print("RAÍZES IMAGINÁRIAS")
	} else if delta > 0{
		fmt.Printf("RAIZES: %.2f e %.2f\n", ((-B - math.Pow(delta, 0.5))/2*A), ((-B + math.Pow(delta, 0.5))/2*A))
		fmt.Print("RAÍZES DISTINTAS")
	} else{
		fmt.Printf("RAIZ: %.2f\n", ((-B - math.Pow(delta, 0.5))/2*A))
		fmt.Print("RAIZ ÚNICA")
	}

}