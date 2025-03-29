package main

import "fmt"

func main (){
	var num int

	fmt.Scan(&num)

	if num > 20 && num < 90{
		fmt.Print("O numero está entre 20 e 90.")
	} else{
		fmt.Print("O numero não está entre 20 e 90.")
	}
}