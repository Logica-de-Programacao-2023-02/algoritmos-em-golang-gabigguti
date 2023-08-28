package main

import "fmt"

func main() {

	var num int
	fmt.Print("Digite um numero de 1 a 7: ")
	fmt.Scan(&num)

	switch num {
	case 1:
		fmt.Print("Domingo")
	case 2:
		fmt.Print("Segunda")
	case 3:
		fmt.Print("Terça")
	case 4:
		fmt.Print("Qurta")
	case 5:
		fmt.Print("Quinta")
	case 6:
		fmt.Print("Sexta")
	case 7:
		fmt.Print("Sabado")
	}
}
