package main

import (
	"fmt"
	"sort"
)

func main() {

	funcionarios := make(map[int]int)
	ordenar := [] int {}
	var meses, idFuncionario int

	for i := 0; i < 100; i++ {
		fmt.Scan(&idFuncionario, &meses)
		if idFuncionario == 0 && meses == 0{
			break
		}

		funcionarios[meses] = idFuncionario
	}

	for indice, _ := range funcionarios{
		ordenar = append(ordenar, indice)
	}

	sort.Ints(ordenar)
	ordenar = append(ordenar[:3])

	for _, valor := range ordenar{
		fmt.Printf("O empregado com ID %d trabalha há %d meses\n", funcionarios[valor], valor)
	}


}