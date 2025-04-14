// Escreva um programa que calcule potências. O usuário deve digitar a base e o expoente, e o programa deve
// apresentar o resultado (sem usar o comando pow). Assuma que o usuário irá digitar valores positivos.

package main

import "fmt"

func main (){
	var  (
		base float64
		expoente int
		resultado float64= 1.0)

	fmt.Print("Insira a base a o expoente, respectivamente, separados por espaço: ")
	fmt.Scan(&base, &expoente)

	for i:= 1; i <= expoente; i++{
		resultado *= base
	}
	
	fmt.Print(resultado)
}