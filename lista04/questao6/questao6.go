package main

import "fmt"

func main (){
	var valores [] int

	for i:= 100; i >= 1; i--{
		valores = append(valores, i)
	}

	fmt.Print(valores)
}