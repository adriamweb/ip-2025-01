package main

import (
    "fmt"
    "math"
)

func main() {
    var volume, areaBase, altura, aresta float64

    fmt.Print("Insira a altura e a aresta da base, respectivamente: ")
    fmt.Scan(&altura, &aresta)

    quadAresta := math.Pow(aresta, 2)
    raizQuad := math.Sqrt(3)

    areaBase = 3.0 / 2.0 * quadAresta * raizQuad
    volume = 1.0 / 3.0 * altura * areaBase

    fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS\n", volume)
}
