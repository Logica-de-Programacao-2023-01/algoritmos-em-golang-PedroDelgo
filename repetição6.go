package main

import "fmt"

func main() {
	var i float64
	var numero float64

	fmt.Print("Digite um numéro: ")
	fmt.Scan(&numero)

	for i = 0; i <= 10; i++ {
		fmt.Println(numero * i)
	}
}
