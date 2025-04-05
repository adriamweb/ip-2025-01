// Uma locadora tem as seguintes regras para aluguel de DVDs:
// - Às segundas, terças e quintas (2, 3 e 5) : desconto de 40% em relação ao preço normal;
// - Às quartas , sextas, sábados e domingos (4,6 ,7 e 1): preço normal;
// - Aluguel de DVDs comuns: preço normal;
// - Aluguel de lançamentos: acréscimo de 15% em relação ao preço normal.
// Desenvolver um programa para ler o preço normal do DVD alugado (em R$) e sua categoria (comum ou
// lançamento). Calcular e imprimir o preço final que será pago pela locação do DVD.

package main

import ("fmt"
		"strings")

func main (){
	var (preco float64
		formaAluguel string)

	fmt.Print("Insira o preço normal do DVD: ")
	fmt.Scan(&preco)

	if preco <= 0{
		fmt.Print("PREÇO INVÁLIDO.")
		return
	}

	fmt.Print("\nC - Preço comum\nL - Preço de lançamento\nInsira a letra associada ao tipo de dvd a ser alugado: ")
	fmt.Scan(&formaAluguel)
	formaAluguel = strings.ToUpper(formaAluguel)

	if formaAluguel != "C" && formaAluguel != "L"{
		fmt.Print("FORMA DE ALUGUEL INVÁLIDA.")
		return
	} else if formaAluguel == "C"{
		fmt.Printf("Às segundas, terças e quintas (2, 3 e 5), o preço será: %.2f\n", preco*0.6)
		fmt.Printf("Às quartas , sextas, sábados e domingos (4,6 ,7 e 1), o preço será: %.2f", preco)
	} else{
		fmt.Printf("Às segundas, terças e quintas (2, 3 e 5), o preço será: %.2f\n", preco * 1.15 - (preco*0.4))
		fmt.Printf("Às quartas , sextas, sábados e domingos (4,6 ,7 e 1), o preço será: %.2f", preco*1.15)
	}

	fmt.Print("")

	
}