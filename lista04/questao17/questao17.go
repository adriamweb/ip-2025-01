package main

import ("fmt"
		"math")

func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    if n == 2 {
        return true
    }
    if n%2 == 0 {
        return false
    }
    sqrt := int(math.Sqrt(float64(n)))
    for i := 3; i <= sqrt; i += 2 {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func main (){

	var vetor [10] int

	for i:= 0; i < 10; i++{
		fmt.Scan(&vetor[i])
	}

	for i, v:= range vetor{
		if isPrime(v){
			fmt.Printf("%d é primo e se encontra na posição %d\n", v, i)
		}
	}
}

