package main

import (
	"fmt"
	"sort"
)

func main (){
	jogadas := make(map[int]int)

	for i:= 0; i<10; i++{
		var numero int
		fmt.Scan(&numero)
		jogadas[numero] ++
	}

	organizar:= make([]int, 0, len(jogadas))
	for i := range jogadas{
		organizar = append(organizar, i)
	}

	sort.Ints(organizar)

	for _, v := range organizar{
		fmt.Printf("O numero %d aparece %d vezes\n", v, jogadas[v])
	}
}