package main

import (
	"fmt"
	"sort"
)

type Valores struct{
	Chave int
	Valor int
}

func main() {
    var idade int
    moda := make(map[int]int)

    for i := 0; i < 10; i++ {
        fmt.Scan(&idade)
        moda[idade]++
    }

	Pares := make([]Valores, 0, len(moda))

	for i, v:= range moda{
		Pares = append(Pares, Valores {Chave : i, Valor :v})
	}

	sort.Slice(Pares, func (i, j int) bool{

		if Pares[i].Valor == Pares[j].Valor {
            return Pares[i].Chave < Pares[j].Chave
        }
		return Pares[i].Valor > Pares[j].Valor
	})

	valorModa := Pares[0].Valor

	for _, v := range Pares{
		if v.Valor == valorModa{
			fmt.Print(v.Chave, " ")
		}
	}
}
