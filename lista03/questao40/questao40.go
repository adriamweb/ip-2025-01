package main

import "fmt"

func main (){
	var P_inicial, D, dP, p_min, lucro, lucroMax, prexoMax float64
	var Q_inicial, dQ, ingressoMax int

	P_inicial = 6
	dP = 0.6
	dQ = 30
	Q_inicial = 130
	D = 300
	p_min = 1

	fmt.Println("Preco Ingressos Vendidos Lucro")
	for P_inicial >= p_min{
		lucro = (P_inicial*float64(Q_inicial))-D
		if lucroMax < lucro{
			lucroMax = lucro
			ingressoMax = Q_inicial
			prexoMax = P_inicial
		}

		fmt.Printf("%.2f ", P_inicial)
		fmt.Printf("%.2d ", Q_inicial)
		fmt.Printf("%.2f\n", lucro)
		P_inicial-= dP
		Q_inicial+= dQ

	}
	fmt.Printf("Lucro maximo: %.2f\n",lucroMax)
	fmt.Printf("Na faixa de preco: %.2f com %.2d ingressos.", prexoMax, ingressoMax)
}