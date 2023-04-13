package main

import "fmt"

func main() {
	var num int
	var atd int

	for {
		fmt.Print("Digite numeros inteiros (para encerrar digite 0): ")
		fmt.Scan(&num)
		if num == 0 {
			break
		}
	}

}
