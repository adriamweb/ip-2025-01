package main

import "fmt"

func main (){
	var local, idaVolta int

	fmt.Print("1 - Região Norte\n2 - Região Nordeste\n3 - Região Centro-Oeste\n4 - Região Sul\nInsira o numero referente ao local de viagem: ")
	fmt.Scan(&local)

	fmt.Print("1 - Viagem com retorno\n2 - Viagem sem retorno\nInsira o numero referente à forma de viagem: ")
	fmt.Scan(&idaVolta)

	if idaVolta == 1{
		if local == 1{
			fmt.Print("O valor da passagem é 900.00.")
		} else if local == 2{
			fmt.Print("O valor da passagem é 650.00.")
		} else if local == 3{
			fmt.Print("O valor da passagem é 600.00.")
		} else if local == 4{
			fmt.Print("O valor da passagem é 550.00.")
		}
	} else{
		if local == 1{
			fmt.Print("O valor da passagem é 500.00.")
		} else if local == 2{
			fmt.Print("O valor da passagem é 350.00.")
		} else if local == 3{
			fmt.Print("O valor da passagem é 350.00.")
		} else if local == 4{
			fmt.Print("O valor da passagem é 300.00.")
		}
	}
}