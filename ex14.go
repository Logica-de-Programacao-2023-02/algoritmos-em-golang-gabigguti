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

	if IMC > 18.5 && peso < 24.9 {
		fmt.Print("Peso normal e seu IMC é ", IMC)
	}
	else if IMC < 18.5 {
		fmt.Print("Peso abaixo do ideal e seu IMC é ", IMC)
	}
	else {
		fmt.Print("Peso acima do ideal e seu IMC é ", IMC)
	}
}