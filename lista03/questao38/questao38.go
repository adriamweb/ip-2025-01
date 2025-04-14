package main

import (
	"fmt"
	"strconv"
)

func main() {
	var cpfEntrada string
	fmt.Scan(&cpfEntrada)

	if len(cpfEntrada) != 11 {
		fmt.Println("CPF inválido: deve conter 11 dígitos")
		return
	}

	cpf := make([]int, 11)
	for i := 0; i < 11; i++ {
		digito, err := strconv.Atoi(string(cpfEntrada[i]))
		if err != nil {
			fmt.Println("CPF inválido: contém caracteres não numéricos")
			return
		}
		cpf[i] = digito
	}

	soma := 0
	for i := 0; i < 9; i++ {
		soma += cpf[i] * (10 - i)
	}
	
	d1 := 0
	if soma%11 >= 2 {
		d1 = 11 - (soma % 11)
	}

	soma = 0
	for i := 0; i < 9; i++ {
		soma += cpf[i] * (11 - i)
	}
	soma += d1 * 2
	d2 := 0
	if soma%11 >= 2 {
		d2 = 11 - (soma % 11)
	}

	if d1 == cpf[9] && d2 == cpf[10] {
		fmt.Println("CPF válido")
	} else {
		fmt.Println("CPF inválido")
	}
}
