package main

import "fmt"

func main (){
	var N1, N2, contador int

	fmt.Print("Insira os dois numeros: ")
	fmt.Scan(&N1, &N2)

	for N1 >= N2{
		contador += 1
		N1-=N2
	}

	fmt.Println("Quociente:", contador)
	fmt.Println("Resto:", N1)
}