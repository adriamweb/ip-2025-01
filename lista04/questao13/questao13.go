package main

import "fmt"

// func ordenarValores(valores[3][2] int) [3][2]int{
// 	var menor, maior, intermediario [2]int

// 	if valores[0][1] > valores[1][1] && valores[0][1] > valores[2][1]{
// 		maior = valores[0]
// 	} else if valores[1][1] > valores[0][1] && valores[1][1] > valores[2][1] {
// 		maior = valores[1]
// 	} else{
// 		maior = valores[2]
// 	}

// 	if valores[0][1] < valores[1][1] && valores[0][1] < valores[2][1]{
// 		menor = valores[0]
// 	} else if valores[1][1] < valores[0][1] && valores[1][1] < valores[2][1] {
// 		menor = valores[1]
// 	} else{
// 		menor = valores[2]
// 	}

// 	intermediario = 
// }

func maiorValor(valores [3][2]int) int {

	if valores[0][1] > valores[1][1] && valores[0][1] > valores[2][1] {
		return 0
	} else if valores[1][1] > valores[0][1] && valores[1][1] > valores[2][1] {
		return 1
	} else {
		return 2
	}
}

func tresMenores(valores [][2]int) [3][2]int {
	var menores [3][2]int

	menores[0] = valores[0]
	menores[1] = valores[1]
	menores[2] = valores[2]

	for i := 3; i < len(valores); i++ {

		if valores[i][1] < menores[0][1] || valores[i][1] < menores[1][1] || valores[i][1] < menores[2][1] {
			menores[maiorValor(menores)] = valores[i]
		}
	}

	return menores
}

func empregadosRecentes() [][2]int {
	var (
		numeroEmpregado, mesesEmpregado int
		empregados                      [][2]int
	)

	for i := 0; i <= 100; i++ {
		fmt.Scan(&numeroEmpregado, &mesesEmpregado)

		if numeroEmpregado == 0 && mesesEmpregado == 0 {
			return empregados
		}

		empregados = append(empregados, [2]int{numeroEmpregado, mesesEmpregado})
	}

	return empregados
}

func show(vetor [3][2]int) {
	fmt.Println("Os empregados mais recentes possuem: ")
	for i := 0; i < len(vetor); i++ {
		fmt.Printf("Numero: %d e trabalha a %d meses.\n", vetor[i][0], vetor[i][1])
	}
}

func main() {
	empregados := empregadosRecentes()

	mostrar := tresMenores(empregados)

	show(mostrar)
}