package main

import "fmt"

func main (){
	var nota1, nota2, nota3, media float32

	fmt.Print("Insira as 3 notas do aluno: ")
	fmt.Scan(&nota1, &nota2, &nota3)

	media = (nota1 + nota2 + nota3)/3
	fmt.Println("MÉDIA:",media)

	if media >= 6 {
		fmt.Print("APROVADO")
	} else {
		fmt.Print("REPROVADO")
	}
		
}