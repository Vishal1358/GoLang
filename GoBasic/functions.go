package main

import "fmt"

func add(x, y int) int {
	var out = x + y
	return out
}

func main() {

	num1 := 5
	num2 := 7
	num3 := 6
	num4 := 6
	num5 := 7
	num6 := 8

	result := add(num1, num2)
	result1 := add(num3, num4)
	result2 := add(num5, num6)
	fmt.Println(result, result1, result2)
}
