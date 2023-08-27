package main

import "fmt"

func main() {

	var valor float32
	fmt.Print("Qual o valor do produto? ")
	fmt.Scan(&valor)

	var desconto float32
	desconto = valor * 0.9
	fmt.Print("O seu valor do produto com o desconto será de ", desconto)
}
