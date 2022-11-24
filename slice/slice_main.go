package main

import "fmt"

func main() {
	var cities []string
	fmt.Println("cities is equal to nil", cities == nil)
	fmt.Printf("cities %#v\n", cities)
	fmt.Println(len(cities))

	number := []int{2, 3, 4, 5}
	fmt.Println(number)

	nums := make([]int, 2)
	fmt.Printf("%#v\n", nums)

	for index, value := range number {
		fmt.Printf("Index: %v, value: %v\n", index, value)
	}

	var n []int
	n = number
	fmt.Println(n)
}
