package main

import "fmt"

func main() {

	var salario, novo float32
	fmt.Print("Qual o seu salário? ")
	fmt.Scan(&salario)

	if salario < 1000 {
		novo = salario * 1.1
		fmt.Print("Seu novo salario será de: ", novo)

	} else if salario >= 1000 {
		novo = salario * 1.05
		fmt.Print("Seu novo salario será de: ", novo)
	}

}
