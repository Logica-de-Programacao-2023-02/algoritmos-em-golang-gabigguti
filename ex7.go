package main

import "fmt"

func main() {

	var numero int
	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	var antecessor int
	antecessor = numero - 1
	fmt.Print("O antecessor desse numero é: ", antecessor)

	var sucessor int
	sucessor = numero + 1
	fmt.Print("O sucessor desse numero é: ", sucessor)
}
