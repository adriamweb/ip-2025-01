package main

import (
	"fmt"
	"strings"
)

func main (){
	var(
		conta, valorPago float64 
		tipoConta string
	)

	fmt.Print("R - Residencial\nC - Comercial\nI - Industrial\nInsira a letra correspondente ao tipo de consumidor: ")
	fmt.Scan(&tipoConta)
	tipoConta = strings.ToUpper(tipoConta)

	fmt.Print("Insira o valor da conta: ")
	fmt.Scan(&conta)

	if conta < 0{
		fmt.Print("VALOR DE CONTA INVÁLIDO")
		return
	}

	if tipoConta == "R"{
		valorPago = 5 + 0.05 * conta
	} else if tipoConta == "C"{
		valorPago = 500
		if conta > 80{
			valorPago += (conta-80)*0.04
		}
	} else if tipoConta == "I"{
		valorPago = 800
		if conta > 100{
			valorPago += (conta - 100)*0.04
		}
	} else{
		fmt.Print("TIPO DE CONSUMIDOR INVÁLIDO")
		return
	}

	fmt.Printf("O valor a ser pago é: %.2f", valorPago)
}