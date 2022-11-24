package main

import "fmt"

func insertionsort(a []int) {
	for i := 1; i < len(a); i++ {
		cvalue := a[i]
		position := i
		for position > 0 && a[position-1] > cvalue {
			a[position] = a[position-1]
			position = position - 1
		}
		a[position] = cvalue
	}
}

func main() {
	a := []int{5, 6, 2, 8, 4, 3}
	fmt.Println("Before sort", a)
	insertionsort(a)
	fmt.Println("After sort", a)

}
