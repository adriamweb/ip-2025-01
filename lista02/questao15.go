// Faça um programa que leia uma data (dia, mês e ano, em uma variável inteira cada), e escreva a mesma data
// no formato dia de (mês por extenso) de ano.

package main

import "fmt"

func main (){
	var dia, mes, ano int

	fmt.Print("Insira, separado por espaços, dia, mês e ano: ")
	fmt.Scan(&dia, &mes, &ano)

	fmt.Print(dia, " ")
	if mes == 1{
		fmt.Print("de Janeiro ")
	} else if mes == 2{
		fmt.Print("de Fevereiro ")
	} else if mes == 3{
		fmt.Print("de Março ")
	} else if mes == 4{
		fmt.Print("de Abril ")
	} else if mes == 5{
		fmt.Print("de Maio ")
	} else if mes == 6{
		fmt.Print("de Junho ")
	} else if mes == 7{
		fmt.Print("de Julho ")
	} else if mes == 8{
		fmt.Print("de Agosto ")
	} else if mes == 9{
		fmt.Print("de Setembro ")
	} else if mes == 10{
		fmt.Print("de Outubro ")
	} else if mes == 11{
		fmt.Print("de Novembro ")
	} else if mes == 12{
		fmt.Print("de Dezembro ")
	} else{
		fmt.Print("INVÁLIDO")
	}

	fmt.Print(ano)
}