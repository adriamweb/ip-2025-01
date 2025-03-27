package main

import "fmt"

func main() {

	var (
		casoTeste int
		qtdPessoas, popular, geral, arquibancada, cadeiras, rendaTotal float64
		rendaFinal[]float64

	)

	fmt.Print("Insira a quantidade de casos teste: ")
	fmt.Scan(&casoTeste)
	

	for i := 1; i <= casoTeste; i = i+1 {

		fmt.Printf("Insira os dados relacionados à arrecadação do jogo %d: ", i)
		fmt.Scan(&qtdPessoas, &popular, &geral, &arquibancada, &cadeiras)

		rendaTotal = qtdPessoas*popular/100 + qtdPessoas*geral/20 + qtdPessoas*arquibancada/10 + qtdPessoas*cadeiras/5
		rendaFinal = append(rendaFinal, rendaTotal)

	}

	for i:= 0; i < len(rendaFinal); i++{
		fmt.Printf("\nA RENDA DO JOGO N° %d E = %.2f", (i + 1),rendaFinal[i] )
	}

}