package main

import "fmt"

func main() {
	a := 5

	for a <= 100 {
		fmt.Println(a + 5)
		a += 5
	}
}
