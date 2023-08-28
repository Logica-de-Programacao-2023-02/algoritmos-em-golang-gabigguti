package main

import "fmt"

func main() {
	var num1, num2, resultado, result int
	fmt.Print("Qual o numero 1? ")
	fmt.Scan(&num1)
	fmt.Print("Qual o numero 2? ")
	fmt.Scan(&num2)

	if num1 > 0 && num2 > 0 {
		resultado = num1 * num2
		fmt.Print("O resultado é: ", resultado)
	} else {
		result = num1 + num2
		fmt.Print("O resultado é: ", result)
	}
}
