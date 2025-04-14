package main

import "fmt"

func main() {

	// começa em 29 e perde 2

	var (
		S         float64
		passo     float64 = 31
		numerador float64 = 1.0
		contador  int     = 1
	)

	for i := 225.0; passo > 1; i -= passo {

		if contador%2 != 0 {
			S += (numerador) / i
		} else {
			S -= (numerador) / i
		}

		numerador *= 2
		contador++
		passo -= 2

	}
	fmt.Print(S)
}
