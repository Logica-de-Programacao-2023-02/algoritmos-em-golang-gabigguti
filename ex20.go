package main

import "fmt"

func main() {

	var idade int
	fmt.Print("Qual a sua idade? ")
	fmt.Scan(&idade)

	if idade < 9 {
		fmt.Print("Classificação: mirim ")
	} else if idade > 10 && idade < 13 {
		fmt.Print("Classificação: infantil")
	} else if idade > 14 && idade < 17 {
		fmt.Print("Classificação: juvenil")
	} else {
		fmt.Print("Classificação: adulto")
	}
}
