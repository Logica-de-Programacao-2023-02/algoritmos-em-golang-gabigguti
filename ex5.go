package main

import "fmt"

func main() {

	var idade int
	fmt.Print("Qual a sua idade? ")
	fmt.Scan(&idade)

	var dias int
	dias = idade * 365
	fmt.Print("A sua idade em dias é: ", dias)

}
