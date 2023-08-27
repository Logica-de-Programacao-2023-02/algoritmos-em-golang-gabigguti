package main

import "fmt"

func main() {

	var dias int
	fmt.Println("Quantos dias você irá trabalhar? ")
	fmt.Scan(&dias)

	var diaria int
	fmt.Println("Qual será o valor da diária? ")
	fmt.Scan(&dias)

	var salario int
	salario = dias * diaria
	fmt.Print("O seu salário será de: ", salario)
}
