package main

import "fmt"

func main () {

	var num1, num2, num3 int
	fmt.Print("Qual o numero 1? ")
	fmt.Scan(&num1)
	fmt.Print("Qual o numero 2? ")
	fmt.Scan(&num2)
	fmt.Print("Qual o numero 3? ")
	fmt.Scan(&num3)

	if num1 > num2 && num1 > num3 {
		fmt.Print("O numero 1 é o maior.")

	} else if num2 > num1 && num2 > num3 {
		fmt.Print("O numero 2 é o maior.")
	} else {
		fmt.Print("O numero 3 é o maior.")
	}
}