package main

import "fmt"

func main() {

	var num int
	fmt.Print("Qual é o número? ")
	fmt.Scan(&num)

	var dobro int
	dobro = num * 2

	var triplo int
	triplo = num * 3

	var quadruplo int
	quadruplo = num * 4

	fmt.Print("O dobro do número é: ", dobro)
	fmt.Print("O triplo do número é: ", triplo)
	fmt.Print("O quadruplo do número é: ", quadruplo)

}
