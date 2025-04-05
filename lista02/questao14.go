package main

import "fmt"

func main (){
	var preco float64
	var vetor[4] int

	fmt.Print("Insira o preço inicial do carro: ")
	fmt.Scan(&preco)

	fmt.Print("Opções de adicionais:\n1- Ar condicionado (R$ 1750,00)\n2- Pintura metálica (R$ 800,00)\n2- Vidro elétrico (R$ 1200,00)\n4- Direção hidráulica (R$ 2000,00)\nInsira 4 valores referentes aos numeros dos adicionais. (Caso nao queira algum dos adicionais, insira 0): ")
	fmt.Scan(&vetor[0], &vetor[1], &vetor[2], &vetor[3])

	for _, valor := range vetor{
		if valor != 0{
			if valor == 1{
				preco += 1750.00
			} else if valor == 2{
				preco += 800.00
			} else if valor == 3{
				preco += 1200.00
			} else if valor == 4{
				preco += 2000.00
			}
		}
	}

	fmt.Printf("O valor final é: %2.f", preco)
}