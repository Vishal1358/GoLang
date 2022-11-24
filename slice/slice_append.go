package main

import "fmt"

func main() {
	number := []int{2, 3, 4, 5}
	number = append(number, 10)
	fmt.Println(number)

	number = append(number, 20, 30, 40, 50)
	fmt.Println(number)

	n := []int{200, 300, 400}
	number = append(number, n...)
	fmt.Println(number)

	// copying element

	src := []int{10, 20, 30}
	// dst := make([]int, len(src))
	// it will copy first 2 element
	dst := make([]int, 2)
	nn := copy(dst, src)
	fmt.Println(src, dst, nn)

	number = append(number[:2], 100)
	fmt.Println(number)

}
