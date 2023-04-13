package main

import "fmt"

func main() {
	var i int
	var num int

	fmt.Print("Digite um numéro: ")
	fmt.Scan(&num)

	for i = 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Println(i)
		}
	}
}
