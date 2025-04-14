// Escreva um programa que receba a idade, a altura e o peso de várias pessoas. Calcule e imprima:
// - a quantidade de pessoas com idade superior a 50 anos;
// - a média das alturas das pessoas com idade entre 10 e 20 anos;
// - a porcentagem de pessoas com peso inferior a 40 quilos entre todas as pessoas analisadas.
// Considere que os dados informados são válidos. Pergunte ao usuário se ele deseja continuar digitando dados ou
// não (Exemplo: 1 - Sim, Outro valor diferente de 1 - Não).

package main

import "fmt"

func criarStruct(age int, height float64, weight float64) Pessoa {
	return Pessoa{idade: age, altura: height, peso: weight}
}

type Pessoa struct {
	idade  int
	altura float64
	peso   float64
}

func main() {
	var idadeT, qtdIdade, contador int
	var pesoT, alturaT, somaAltura, qtdPeso float64
	var pessoas []Pessoa

	for i := 1; i == 1; {
		fmt.Print("Insira a idade, altura e peso da pessoa, respectivamente: ")
		fmt.Scan(&idadeT, &alturaT, &pesoT)
		pessoa := criarStruct(idadeT, alturaT, pesoT)

		pessoas = append(pessoas, pessoa)

		fmt.Print("1 - Para continuar\n? - Outro valor para parar\nInsira o numero:")
		fmt.Scan(&i)

	}

	for _, valor := range pessoas {
		if valor.idade > 50 {
			qtdIdade++
		}

		if valor.peso < 40 {
			qtdPeso++
		}

		if valor.idade < 20 && valor.idade > 10 {
			somaAltura += valor.altura
			contador++
		}

	}

	fmt.Println("Quantidade de pessoas com idade maior que 50:", qtdIdade)
	fmt.Println("Média das alturas da pessoa com idade entre 10 e 20 anos: ", somaAltura/float64(contador))
	fmt.Println("Porcentagem de pessoas com peso inferior a 40 kilos: ", qtdPeso/float64(len(pessoas)))
}
