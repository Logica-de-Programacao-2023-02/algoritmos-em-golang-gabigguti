package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	if num%3 == 0 && num%5 == 0 {
		fmt.Print("É divisivel por 3 e 5 ao mesmo tempo.")
	} else {
		fmt.Print("Não é divisível opr 3 e 5 ao mesmo tempo.")
	}
}
