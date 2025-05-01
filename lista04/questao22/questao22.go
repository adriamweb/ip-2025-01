package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func clearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func flushStdin() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func waitForEnter() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n\nPressione Enter para voltar ao menu principal...")
	reader.ReadString('\n')
}

func menu() int {
	var numero int

	clearTerminal()

	fmt.Print(strings.Repeat("-", 25))
	fmt.Print("\n1. Efetuar depósito\n2. Efetuar saque\n3. Consultar saldo\n4. Finalizar o programa\n")
	fmt.Print(strings.Repeat("-", 25))
	fmt.Print("\nInsira o número referente à operação a ser realizada: ")

	for {
		fmt.Scan(&numero)
		flushStdin()

		if numero > 4 || numero < 1 {
			fmt.Print("\nNúmero inválido. Insira novamente: ")
			continue
		}
		break
	}

	clearTerminal()
	return numero
}

func verify(codigo int, codigos [10]int) bool {
	for _, v := range codigos {
		if v == codigo {
			return true
		}
	}
	return false
}

func returnPosition(codigo int, codigos [10]int) int {
	for i, v := range codigos {
		if codigo == v {
			return i
		}
	}
	return 0
}

func ativoBancario(saldos [10]float64) float64 {
	var soma float64
	for _, v := range saldos {
		soma += v
	}
	return soma
}

func main() {
	var (
		codigo   int
		contador int
		codigos  [10]int
		saldos   [10]float64
	)

	for contador < 3 {
		fmt.Printf("\nInsira o código da %d° conta: ", contador+1)

		for {
			fmt.Scan(&codigo)
			flushStdin()

			if verify(codigo, codigos) {
				fmt.Print("\nO código informado já está associado a uma conta. Insira outro: ")
				continue
			}

			codigos[contador] = codigo
			break
		}

		fmt.Printf("\nInsira o saldo da %d° conta: ", contador+1)
		fmt.Scan(&saldos[contador])
		flushStdin()
		contador++
	}

	for {
		var codigo int
		numero := menu()

		switch numero {
		case 1:
			fmt.Print("Insira o código da conta: ")
			for {
				fmt.Scan(&codigo)
				flushStdin()

				if !verify(codigo, codigos) {
					fmt.Print("\nO código inserido não está associado a nenhuma conta.")
					waitForEnter()
					clearTerminal()
					break
				} else {
					var deposito float64

					fmt.Print("\nInsira o valor do depósito a ser feito: ")
					fmt.Scan(&deposito)
					flushStdin()

					saldos[returnPosition(codigo, codigos)] += deposito
					fmt.Print("\nDepósito realizado com sucesso!")
					waitForEnter()
					clearTerminal()
					break
				}
			}

		case 2:
			fmt.Print("Insira o código da conta: ")
			for {
				fmt.Scan(&codigo)
				flushStdin()

				if !verify(codigo, codigos) {
					fmt.Print("\nO código inserido não está associado a nenhuma conta.")
					waitForEnter()
					clearTerminal()
					break
				} else {
					var saque float64

					fmt.Print("\nInsira o valor do saque a ser feito: ")
					fmt.Scan(&saque)
					flushStdin()

					saldos[returnPosition(codigo, codigos)] -= saque
					fmt.Print("\nSaque realizado com sucesso!")
					waitForEnter()
					clearTerminal()
					break
				}
			}

		case 3:
			fmt.Printf("\nValor referente ao ativo bancário: %.2f", ativoBancario(saldos))
			fmt.Print("\nConsulta realizada com sucesso!")
			waitForEnter()
			clearTerminal()

		case 4:
			fmt.Print("Programa encerrado...")
			return
		}
	}
}
