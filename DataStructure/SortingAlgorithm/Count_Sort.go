package main

import (
	"fmt"
)

func findMaxElement(A []int) int {
	max_num := A[0]
	for i := 0; i < len(A); i++ {
		if A[i] > max_num {
			max_num = A[i]
		}
	}
	return max_num
}
func countsort(A []int) {
	n := len(A)
	maxsize := findMaxElement(A)
	carray := make([]int, (maxsize + 1))
	for i := 0; i < n; i++ {
		carray[A[i]] = carray[A[i]] + 1
	}
	i := 0
	j := 0
	for i = 0; i < maxsize+1; {
		if carray[i] > 0 {
			A[j] = i
			j++
			carray[i] = carray[i] - 1
		} else {
			i++
		}
	}
}

func main() {
	A := []int{5, 9, 9, 6, 3, 2, 7, 5, 4, 8}
	fmt.Println("Original Array:", A)
	countsort(A)
	fmt.Println("Sorted Array:", A)

}
