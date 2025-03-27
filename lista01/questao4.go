package main

import "fmt"

func main (){
	var salario, qtdKw, valorUnitKw, valorFinal, valorDesconto float64 

	fmt.Print("Insira o valor do salario minimo: ")
	fmt.Scan(&salario)

	fmt.Print("Insira a quantidade de Kw gasto pela residencia: ")
	fmt.Scan(&qtdKw)

	valorUnitKw = (7.0/1000.0) * salario
	valorFinal = valorUnitKw * qtdKw
	valorDesconto = valorFinal * 0.9

	fmt.Printf("\nCusto por Kw: %.2f\n", valorUnitKw)
	fmt.Printf("Custo do consumo: %.2f\n",valorFinal)
	fmt.Printf("Custo com desconto: %.2f\n", valorDesconto)

}