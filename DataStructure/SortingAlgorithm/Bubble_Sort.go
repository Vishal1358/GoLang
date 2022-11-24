package main

import "fmt"

func BubbleSort(a []int) {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
	}
}

func main() {
	a := []int{2, 5, 7, 3, 9, 6}
	fmt.Println("Before Sort", a)
	BubbleSort(a)
	fmt.Println("Before Sort", a)
}
