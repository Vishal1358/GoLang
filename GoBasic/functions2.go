package main

import "fmt"

func calc(x, y int) (int, int) {
	var out1 = x + y
	var out2 = x - y
	return out1, out2
}

func main() {

	num1 := 8
	num2 := 7

	result1, result2 := calc(num1, num2)
	// result1 := add(num3, num4)
	// result2 := add(num5, num6)
	fmt.Println(result1, result2)
}
