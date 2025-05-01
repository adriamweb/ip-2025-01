// Faça um programa que leia um primeiro vetor com dez números inteiros e um segundo vetor com cinco
// números inteiros. Mostre uma lista dos números do primeiro vetor com seus respectivos divisores armazenados
// no segundo vetor, bem como suas posições.

package main

import "fmt"

func main() {
	var (
		vetor1 = [10]int{}
		vetor2 = [5]int{}
	)

	for i := 0; i < 10; i++ {
		fmt.Scan(&vetor1[i])
	}

	for i := 0; i < 5; i++ {
		fmt.Scan(&vetor2[i])
	}

	for _, v := range vetor1 {
		
		for j, z := range vetor2 {
			if v % z == 0{
				fmt.Printf("%d divisível por %d na posição %d\n", v, z, j)
			}
		}
	
	}
}
