package main

import "fmt"

func main (){
	var compra float64

	fmt.Print("Insira o valor da compra: ")
	fmt.Scan(&compra)

	if 0 < compra && compra < 10.0{
		compra *=1.7
		fmt.Printf("O valor da venda é: %.2f",compra)
	} else if 10.0 <=compra && compra < 30.0{
		compra *= 1.5
		fmt.Printf("O valor da venda é: %.2f",compra)
	} else if 30.0 <= compra && compra < 50.0{
		compra *= 1.4
		fmt.Printf("O valor da venda é: %.2f",compra)
	} else if 50 <= compra{
		compra *= 1.3
		fmt.Printf("O valor da venda é: %.2f",compra)
	} else{
		fmt.Print("Valor de compra inválido.")
	}
}