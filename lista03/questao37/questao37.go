package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var numero, teste, contador, soma uint64

	fmt.Print("Insira o numero: ")
	fmt.Scan(&numero)
	variavel := strconv.Itoa(int(numero))
	iStr := "1" + strings.Repeat("0", len(variavel)-1)
	contador = uint64(len(variavel))

	i, err := strconv.ParseUint(iStr, 10, 64)
	if err != nil {
		fmt.Print("Erro!")
		return
	}

	for i >= 1 {
		teste = numero / i
		numero -= teste * i
		i /= 10
		soma += teste * uint64(math.Pow(8, float64(contador)-1))
		contador--
	}

	fmt.Print("O numero em base decimal é: ", soma)
}
