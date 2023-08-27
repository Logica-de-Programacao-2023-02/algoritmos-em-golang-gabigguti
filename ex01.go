package main

import "fmt"

func main() {

	var numero1 int
	fmt.Print("Qual é o primeiro número? ")
	fmt.Scan(&numero1)

	var numero2 int
	fmt.Print("Qual é o segundo numero? ")
	fmt.Scan(&numero2)

	var numero3 int
	fmt.Print("Qual é o segundo numero? ")
	fmt.Scan(&numero3)

	var resultado int
	resultado = numero1 + numero2 + numero3
	fmt.Print("O resultado é ", resultado)

}
