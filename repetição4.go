package main

import "fmt"

func main() {
	var i uint

	for i = 0; i <= 30; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
