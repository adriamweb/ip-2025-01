package main

import (
	"fmt"
	"math"
	"strings"
)

func main (){
	fmt.Printf("%-10s| %-10s\n", "Raio", "Volume")
	linha:= strings.Repeat("-", 30)
	fmt.Println(linha)

	 for i:= 0.0; i<= 20; i+= 0.5{
		fmt.Printf("%-10.2f| %-10.5f\n", i, math.Pi*4.0/3.0*math.Pow(i,3))

	 }
}