package main

import "fmt"

func main (){
	var matricula, horaExtra int
	const salarioMinimo float64 = 788.00
	const valorHoraExtra float64 = 10.00
	var salarioBruto float64 = 3*salarioMinimo
	fmt.Print(salarioBruto,"\n")
	var salarioLiquido float64

	fmt.Print("Insira a matricula do funcionario: ")
	fmt.Scan(&matricula)

	fmt.Print("Insira a quantidade de horas extras do funcionario: ")
	fmt.Scan(&horaExtra)

	salarioBruto += valorHoraExtra * float64(horaExtra)

	if salarioBruto > 1500{
		salarioLiquido += salarioBruto*0.88
	}

	if salarioBruto > 2000{
		salarioLiquido -= salarioBruto*0.2
	}

	fmt.Printf("A matrícula do funcionário é: %d\n",matricula)
	fmt.Printf("O salário bruto é: %.2f\n", salarioBruto)
	fmt.Print("O salário liquido é: ", salarioLiquido)
	
}



