package main

import "fmt"

func verify(local [24] int) []int{
	var disponiveis[] int

	for i, v:= range local{
		if v == 0{
			disponiveis = append(disponiveis, i)
		}
	}

	return disponiveis
}

func main (){
	var (
		disponiveis [] int
		janela = [24] int {0, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1}
		corredor = [24] int {0, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1}
		local int
	)
	
	fmt.Print("1. Corredor\n2. Janela\n\nInsira o número referente à localização da poltrona que deseja comprar: ")
	for{

		fmt.Scan(&local)
		if local > 2 || local < 1{
			fmt.Print("Número inválido. Insira novamnete: ")
			continue
		}

		break
	}

	if local == 1{
		disponiveis = verify(corredor)
	} else{
		disponiveis = verify(janela)
	}

	if len(disponiveis) == 0{
		fmt.Print("Não há cadeiras disponíveis para a venda.")
	} else{
		fmt.Print("As cadeiras disponúveis são: ", disponiveis)
	}

	
}