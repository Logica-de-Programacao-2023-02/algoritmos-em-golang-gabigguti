package main

import "fmt"

func main() {

	var num int
	fmt.Print("Digite um numero para descobrir se é par ou impar.")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Print("O número é par")
	} else {
		fmt.Print("O numero é ímpar.")
	}
}
