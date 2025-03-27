package main

import "fmt"

func main(){
	const π = 3.14159
	var raio, altura, Ac, Al, At float32

	fmt.Print("Insira o raio da lata (em M): ")
	fmt.Scan(&raio)

	fmt.Print("Insira a altura da lata (em M): ")
	fmt.Scan(&altura)

	Ac = π *raio * raio
	Al = 2*π*raio*altura
	At = 2*Ac + Al

	fmt.Printf("O VALOR DO CUSTO E: %.2f\n",At*100)




}