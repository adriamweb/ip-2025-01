package main

import (
	"fmt"
)

func Verify (valor int, valores[] int) bool{
	for _, x := range valores{
		if valor == x{
			return true
		}
	}

	return false
}

func Show (x [][] int, qtdNotas int){
	fmt.Println("\nNota   Frequência Relativa   Frequência Absoluta")

	for _, valor := range x{
		fmt.Printf("%3.d %13.d %23.2f\n", valor[0], valor[1],float64(valor[1])/float64(qtdNotas))
	} 
}

func main (){
	var (
		frequencia [][] int
		notas [15] int
		contados[] int
		qtd int
	)

	for i:= 0; i < 15; i++{
		fmt.Scan(&notas[i])
	}

	for _, valor := range notas{
		if !Verify(valor, contados){
			qtd = 0
			for _, j := range notas{

				if j == valor{
					qtd += 1
				}
			}

			frequencia = append(frequencia, []int{valor, qtd})
			contados = append(contados, valor)
		}
	}

	Show(frequencia, 15)
}