package main

import "fmt"

func main() {
	var qtd int
	var num int
	var soma int

	for {
		fmt.Print("digite numeros inteiros (digite 0 para parar): ")
		fmt.Scan(&num)
		if num == 0 {
			break
		}
	}
	soma += num

	fmt.Print("A media dos seus numeros é: ", soma/qtd)

}
