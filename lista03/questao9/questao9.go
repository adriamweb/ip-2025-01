// Faça um programa que receba 2 notas de N alunos. Calcule e mostre:
// a) a média aritmética das 2 notas de cada aluno;
// b) uma mensagem de acordo com as regras a seguir:
// Média Aritmética Mensagem
// Até 3 Reprovado
// Entre 3 e 7 Exame
// De 7 para cima Aprovado
// c) o total de alunos aprovados;
// d) o total de alunos de exame;
// e) o total de alunos reprovados;
// f) a média da classe.
// Assuma que o N informado é válido, assim como as 2 notas de cada aluno. 

package main

import "fmt"

func main (){
	var (
		notas[] float64
		nota1, nota2, soma float64
		qtdAlunos, aprovados, reprovados, exame int
	)

	fmt.Print("Insira a quantidade de alunos que deseja cadastrar a nota: ")
	fmt.Scan(&qtdAlunos)

	for i:= 1; i <= qtdAlunos; i++{
		fmt.Printf("\nInsira as duas notas do %d aluno separadas por um espaço: ", i)
		fmt.Scan(&nota1, &nota2)
		soma += (nota1 + nota2)/2

		notas = append(notas,(nota1 + nota2)/2)

	} 

	for indice, valor:= range notas{
		fmt.Printf("\nInformações referentes ao aluno %d: \n", indice + 1)

		fmt.Println("Média aritmetica: ", valor)
		
		if valor < 3{
			fmt.Println("Reprovado")
			reprovados ++
		} else if valor >= 3 && valor < 7{
			fmt.Println("Exame")
			exame ++
		} else{
			fmt.Println("Aprovado")
			aprovados ++
		}
	}

	fmt.Println("Alunos aprovados: ", aprovados)
	fmt.Println("Alunos em exame: ", exame)
	fmt.Println("Alunos reprovados: ", reprovados)
	fmt.Printf("A média da clase é: %.2f", soma/float64(len(notas)))

}