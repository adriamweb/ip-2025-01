package main

import (
    "fmt"
    "math"
)

func maclaurinSeno(A float64) float64 {
    return A - (math.Pow(A, 3) / 6) + (math.Pow(A, 5) / 120) - (math.Pow(A, 7) / 5040)
}

func main() {
	fmt.Print("Ângulo    Seno\n")
    fmt.Println("-----------------")

    for A := 0.0; A <= 6.3; A += 0.1 {
        senoA := maclaurinSeno(A)
		if senoA < 0{
			fmt.Printf("%.1f  |  %.5f\n", A, senoA)
		} else{
			fmt.Printf("%.1f  |   %.5f\n", A, senoA)
		}
    }
}
