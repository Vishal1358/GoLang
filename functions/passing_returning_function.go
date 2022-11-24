package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func arth(option string) func(int, int) int {
	if option == "+" {
		return add
	} else {
		return sub
	}
}

func main() {
	var op string
	fmt.Scan(&op)
	res := arth(op)
	fmt.Println(res(20, 30))

}
