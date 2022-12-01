package main

import "fmt"

func mergesort(A []int, left int, right int) {
	if left < right {
		mid := (left + right) / 2
		mergesort(A, left, mid)
		mergesort(A, mid+1, right)
		merge(A, left, mid, right)
	}
}

func merge(A []int, left int, mid int, right int) {
	var i, j, k int
	i = left
	j = mid + 1
	k = left
	B := make([]int, (right + 1))

	for i <= mid && j <= right {
		if A[i] < A[j] {
			B[k] = A[i]
			i = i + 1
		} else {
			B[k] = A[j]
			j++
		}
		k++
	}
	for i <= mid {
		B[k] = A[i]
		i++
		k++
	}
	for j <= right {
		B[k] = A[j]
		j++
		k++
	}
	for x := left; x < right+1; x++ {
		A[x] = B[x]
	}
}

func main() {
	A := []int{25, 36, 45, 22, 34, 56, 19}
	fmt.Println("Original List", A)
	mergesort(A, 0, (len(A) - 1))
	fmt.Println("Sorted List", A)
}
