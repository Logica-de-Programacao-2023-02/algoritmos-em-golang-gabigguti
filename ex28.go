package main

import "fmt"

func main() {

	var num int
	fmt.Print("Informe um numero: ")
	fmt.Scan(&num)

	for i := 1; i > 0; i++ {
		if num%i == 0 {
			fmt.Println(i)
		} else {
			continue
		}
	}
}
