package main

import "fmt"

func main() {
	var i uint
	for i = 0; i <= 19; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
