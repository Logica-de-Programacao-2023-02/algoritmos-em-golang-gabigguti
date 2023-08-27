package main

import "fmt"

func main() {

	var num1 int
	fmt.Println("Informe um numero: ")
	fmt.Scan(&num1)

	var num2 int
	fmt.Println("Informe outro numero: ")
	fmt.Scan(&num2)

	var num3 int
	fmt.Println("Informe outro numero: ")
	fmt.Scan(&num3)

	var resultado float32
	resultado = ((num1 * 2) + (num2 * 3) + (num3 * 5)) / 10
	fmt.Print("O resultado é ", resultado)
}
