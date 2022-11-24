package main

import "fmt"

func shellsort(A []int) {
	n := len(A)
	gap := n / 2
	for gap != 0 {
		i := gap
		for i < n {
			temp := A[i]
			j := i - gap
			for j >= 0 && A[j] > temp {
				A[j+gap] = A[j]
				j = j - gap
			}
			A[j+gap] = temp
			i = i + 1
		}
		gap = gap / 2
	}
}

func main() {
	A := []int{3, 5, 2, 7, 4, 1}
	fmt.Println("Before sort", A)
	shellsort(A)
	fmt.Println("after sort", A)

}
