package main

import "fmt"

func main() {

	var altura float32
	fmt.Print("Qual a sua altura? ")
	fmt.Scan(&altura)

	var peso float32
	fmt.Print("Qual o seu peso? ")
	fmt.Scan(&peso)

	var IMC float32
	IMC = peso / (altura * altura)
	fmt.Print("O seu IMC é de: ", IMC)

}
