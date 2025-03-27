package main

import (
	"fmt"
	"strings"
)

func main(){

	var (
		contaCliente int
		consumoAgua, valorConta float32
		tipoConsumidor string
	)

	fmt.Print("Insira a conta do cliente, consumo de água por metros cúbicos e o tipo do consumidor, respectivamente: ")
	fmt.Scan(&contaCliente, &consumoAgua, &tipoConsumidor)
	tipoConsumidor = strings.ToUpper(tipoConsumidor)

	fmt.Println("CONTA =",contaCliente)

	if tipoConsumidor == "R"{
		valorConta = 5 + 0.05 * consumoAgua

		fmt.Print("VALOR DA CONTA: ",valorConta)

	} else if tipoConsumidor == "C"{
		valorConta = 500
		if consumoAgua > 80{
			valorConta = valorConta + (consumoAgua-80)*0.25
		}

		fmt.Print("VALOR DA CONTA: ",valorConta)

	} else if tipoConsumidor == "I"{ 
		valorConta = 800
		if consumoAgua > 100 {
			valorConta = valorConta + (consumoAgua - 100)*0.04
		}

		fmt.Print("VALOR DA CONTA: ",valorConta)

	} else{
		fmt.Print("TIPO DE CONSUMIDOR INVÁLIDO.")
	}


}