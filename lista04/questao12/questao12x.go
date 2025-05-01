package main

import (
	"fmt"
	"sort"
)

func main() {
	dados := make(map[int]int)
	organizar:= []int {}

	var chave int

	for i := 0; i < 15; i++ {
		fmt.Scan(&chave)
		dados[chave] ++
	}

	for indice, _ := range dados{
		organizar = append(organizar, indice)
	}

	sort.Ints(organizar)

	fmt.Println("\nNota   Frequência Relativa   Frequência Absoluta")

	for _, valor := range organizar{
		fmt.Printf("%3.d %13.d %23.2f\n", valor, dados[valor], float64(dados[valor])/float64(15))
	}
}