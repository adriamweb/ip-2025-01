package main

import "fmt"

func main(){

	var a,b,c,d, det float32

	fmt.Print("Insira o primeiro elemento: ")
	fmt.Scan(&a)

	fmt.Print("Insira o segundo elemento: ")
	fmt.Scan(&b)

	fmt.Print("Insira o terceiro elemento: ")
	fmt.Scan(&c)

	fmt.Print("Insira o qaurto elemento: ")
	fmt.Scan(&d)

	var matriz[2][2]float32 = [2][2]float32{{a,b},{c,d}}

	det = matriz[0][0]*matriz[1][1] - matriz[0][1]*matriz[1][0]

	if det > 9 || det < -9{
		fmt.Print("VALOR INVÁLIDO.")
	} else{
		fmt.Printf("O VALOR DO DETERMINANTE E = %.2f", det)
	}
	

}