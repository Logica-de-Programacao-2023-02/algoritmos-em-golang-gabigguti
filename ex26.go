package main

import "fmt"

func main() {

	var num int
	fmt.Print("Informe um numero: ")
	fmt.Scan(&num)

	var mult int = 1

	for i := 0; i < 10; i++ {
		fmt.Println(num * mult)
		mult++

	}
}
