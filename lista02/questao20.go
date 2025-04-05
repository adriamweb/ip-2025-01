package main

import "fmt"

func main (){
	var preco float64
	var etiqueta int

	fmt.Print("Insira o preço: ")
	fmt.Scan(&preco)

	fmt.Print("Insira o código de pagamento (de 1 a 4): ")
	fmt.Scan(&etiqueta)

	if etiqueta == 1 {
		fmt.Printf("O valor a ser pago é: %.2f", preco*0.9)
	} else if etiqueta == 2{
		fmt.Printf("O valor a ser pago é: %.2f", preco*0.95)
	} else if etiqueta == 3{
		fmt.Printf("O valor a ser pago é: %.2f", preco)
	} else if etiqueta == 4{
		fmt.Printf("O valor a ser pago é: %.2f", preco*1.1)
	} else{
		fmt.Print("VALOR DE ETIQUETA INVÁLIDO")
	}
}