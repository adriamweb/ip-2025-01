package main

import "fmt"

func main (){
	var identificador int
	var notas [3]float64
	var mediaExercicios, mediaFinal float64
	var conceito, aprovacao string

	fmt.Print("Insira o numero identificador do aluno: ")
	fmt.Scan(&identificador)

	fmt.Print("Insira as 3 notas do aluno separadas por espaço: ")
	fmt.Scan(&notas[0], &notas[1], &notas[2])

	mediaExercicios = (notas[0] + notas[1] + notas[2])/3
	mediaFinal = ((notas[0] + notas[1]*2 + notas[2]*3 + mediaExercicios)/7)

	if mediaFinal >= 9.00 && mediaFinal <= 10.0 {
		conceito = "A"
		aprovacao = "APROVADO"
	} else if mediaFinal >= 7.5 && mediaFinal < 9.0 {
		conceito = "B"
		aprovacao = "APROVADO"
	} else if mediaFinal >= 6.00 && mediaFinal < 7.5 {
		conceito = "C"
		aprovacao = "APROVADO"
	} else if mediaFinal >= 4.00 && mediaFinal < 6.0 {
		conceito = "D"
		aprovacao = "REPROVADO"
	} else if mediaFinal < 4.0 {
		conceito = "E"
		aprovacao = "REPROVADO"
	} 

	fmt.Print(identificador, " ", notas[0], " ",notas[1], " ",notas[2], " ",mediaExercicios, " ",mediaFinal, " ",conceito," ", aprovacao)

}