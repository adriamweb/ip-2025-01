// Escreva um programa que receba vários números, calcule e mostre:
// a) a soma dos números digitados;
// b) a quantidade de números digitados;
// c) a média dos números digitados;
// d) o maior número digitado;
// e) o menor número digitado;
// f) a média dos números pares;
// g) a percentagem dos números ímpares entre todos os números digitados.
// Finalize a entrada de dados com a digitação do número 30.000.

package main

import "fmt"

func main (){

	var numero, qtdNumeros, soma, qtdImpares, maiorNumero, menorNumero, qtdPares, somaPares int

	for{
		fmt.Printf("\nInsira o %d numero: ", qtdNumeros+1)
		fmt.Scan(&numero)

		if numero == 30000{
			break
		}
		
		qtdNumeros ++
		soma += numero
		
		if numero % 2.0 == 0{
			qtdPares ++
			somaPares += numero
		} else{
			qtdImpares ++
		}

		if maiorNumero < numero{
			maiorNumero = numero
		}
		
		if menorNumero > numero || (menorNumero == 0 && qtdNumeros == 1){
			menorNumero = numero
		}
	}

	if qtdNumeros > 1{

		fmt.Println("\nA soma dos numeros digitados é: ", soma)
		fmt.Println("A quantidade de numeros digitados é: ", qtdNumeros)
		fmt.Println("A média dos numeros digitados é: ", float64(soma)/float64(qtdNumeros))
		fmt.Println("O maior numero digitado é: ", maiorNumero)
		fmt.Println("O menor numero digitado é: ", menorNumero)
		fmt.Println("A média dos numeros pares é: ", float64(somaPares)/float64(qtdPares))
		fmt.Println("A percentagem de numeros impares é: ", float64(qtdImpares)/float64(qtdNumeros))
	} else{
		fmt.Print("Nao existem numeros armazenados.")
	}


}