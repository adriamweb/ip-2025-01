package main

import "fmt"

func main (){
	var(
		centena, dezena, unidade, concatenado, quadrado int
	)

	fmt.Print("Insira um numero para a centena: ")
	fmt.Scan(&centena)

	if centena > 9 || centena < 0{
		fmt.Print("DIGITO INVALIDO\n")
		return
	}

	fmt.Print("Insira um numero para a dezena: ")
	fmt.Scan(&dezena)

	if dezena > 9 || dezena < 0{
		fmt.Print("DIGITO INVALIDO\n")
		return
	}

	fmt.Print("Insira um numero para a unidade: ")
	fmt.Scan(&unidade)

	if unidade > 9 || unidade < 0 {
		fmt.Print("DIGITO INVALIDO\n")
		return
	}

	centena *= 100
	dezena *= 10
	concatenado = centena + dezena + unidade

	quadrado = concatenado*concatenado
	fmt.Print(concatenado, ", ",quadrado)
}