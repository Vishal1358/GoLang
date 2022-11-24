package main

import "fmt"

func main() {
	var a, b = []int{1, 2, 3}, []int{1, 2, 3}

	var eq bool = true
	for i, valueA := range a {
		if valueA != b[i] {
			eq = false
			break
		}

	}
	if eq {
		fmt.Println("a and b slices are equal")
	} else {
		fmt.Println("a and b slices are not equal")
	}
}
