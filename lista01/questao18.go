package main

import "fmt"

func main (){

	var valorInicial, razao, nTermos, soma, ultimoTermo int

	fmt.Print("Insira os 3 valores: ")
	fmt.Scan(&valorInicial, &razao, &nTermos)

	ultimoTermo = valorInicial + (nTermos-1)*razao

	for i:= valorInicial; i != ultimoTermo; i = i + razao{

		soma = soma + i

		if i == ultimoTermo - razao {
			soma = soma+ i + razao
		}
	}

	fmt.Print(soma,"\n")

}