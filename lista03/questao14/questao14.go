package main

import "fmt"

func numeroPrimo(numero int) bool{
	for j:= 2; j < (numero/2);j++{
		if numero % j == 0{
			return false
		}
	}
	return true
}

func main() {
	var n1, n2 int
	var valor bool

	fmt.Print("Insira os dois numeros separados por espao: ")
	fmt.Scan(&n1, &n2)

	for i:= n1; i<=n2; i++{
		valor = numeroPrimo(i)
		if valor{
			fmt.Println(i)
		}
	}
}
