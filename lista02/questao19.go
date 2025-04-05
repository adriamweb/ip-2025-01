// Desenvolver um programa com as opções de calcular e imprimir o volume e a área da superfície de um cone
// reto, de um cilindro ou de uma esfera. O programa deverá ler a opção da figura desejada (1-cone / 2-cilindro /
// 3-esfera) e de acordo com a opção escolhida calcular e escrever o volume e a área da superfície da figura
// pedida. Fórmulas:

package main

import ("fmt"
		"math")

func main (){
	const π float64 = 3.14159
	var figura int
	var raio, altura, area, volume float64

	fmt.Print("Insira o raio: ")
	fmt.Scan(&raio)

	fmt.Print("Insira a altura: ")
	fmt.Scan(&altura)

	fmt.Print("1 - Cone\n2 - Cilindro\n3 - Esfera\nInsira o numero correspondente à figura que quer: ")
	fmt.Scan(&figura)

	if  figura == 1{
		volume= π * math.Pow(raio,2)*altura/3
		area = π * raio * math.Pow((raio*raio + altura * altura), 0.5)

	} else if figura == 2{
		volume = π * math.Pow(raio,2)*altura
		area= 2* π * raio * altura
		
	} else if figura == 3{
		volume = 4/3*π*math.Pow(raio,3)
		area = 4* π * math.Pow(raio,2)
	} else{
		fmt.Print("VALOR INVÁLIDO")
		return
	}

	fmt.Printf("Volume: %.4f", volume)
	fmt.Printf("\nArea: %.4f", area)
}