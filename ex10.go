package main

import "fmt"

func main() {

	var peso float32
	fmt.Println("Qual o seu peso em kg? ")
	fmt.Scan(&peso)

	var libras float32
	libras = peso / 0.45
	fmt.Print("O seu peso em libras é de: ", libras)
}
