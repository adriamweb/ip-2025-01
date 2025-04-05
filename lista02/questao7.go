package main

import "fmt"

func main() {
    var A, B, C, MAIOR, MENOR, INTER int

    fmt.Print("Insira o primeiro número inteiro: ")
    fmt.Scan(&A)

    for {
        fmt.Print("Insira o segundo número inteiro (diferente do primeiro): ")
        fmt.Scan(&B)
        if B != A {
            break
        }
        fmt.Println("Os números não podem ser iguais.")
    }

    for {
        fmt.Print("Insira o terceiro número inteiro (diferente dos dois anteriores): ")
        fmt.Scan(&C)
        if C != A && C != B {
            break
        }
        fmt.Println("Os números não podem ser iguais.")
    }

    if A > B && A > C {
        MAIOR = A
    } else if B > A && B > C {
        MAIOR = B
    } else {
        MAIOR = C
    }

    if A < B && A < C {
        MENOR = A
    } else if B < A && B < C {
        MENOR = B
    } else {
        MENOR = C
    }

    INTER = A + B + C - MAIOR - MENOR

    fmt.Printf("Os números em ordem crescente: %d, %d, %d\n", MENOR, INTER, MAIOR)
}