package main

import "fmt"

func main() {

	var salario float32
	fmt.Print("Qual é o seu salário? ")
	fmt.Scan(&salario)

	var aumento float32
	aumento = salario * (1.15)
	fmt.Print("O seu salário com o aumento será de ", aumento)
}
