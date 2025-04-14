package main

import "fmt"

// Num frigorífico existem 90 bois. Cada boi traz preso no seu pescoço um cartão contendo um número de
// identificação e seu peso. Implementar um programa que escreva o número e o peso do boi mais gordo e do boi
// mais magro (não é necessário armazenar os dados de todos os bois)
func main() {

	var (
		magro, gordo, peso     float64
		idMagro, Idgordo, id int
	)

	fmt.Print("Insira o id e o peso do 1° boi: ")
	fmt.Scan(&Idgordo, &gordo)

	magro = gordo
	idMagro = Idgordo

	for i := 2; i <= 5; i++ {
		fmt.Printf("Insira o id e o peso do %d boi: ", i)
		fmt.Scan(&id, &peso)

		if peso > gordo{
			gordo = peso
			Idgordo = id
		} else if peso < magro{
			magro = peso
			idMagro = id
		}
	}

	fmt.Println("Id e peso do boi mais pesado: ", Idgordo, gordo)
	fmt.Println("Id e peso do boi mais leve: ", idMagro, magro)
}
