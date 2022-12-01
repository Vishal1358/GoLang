package main

import "fmt"

func partition(A []int, low int, high int) int {
	pivot := A[low]
	i := low + 1
	j := high
	for i < j {
		for A[i] <= pivot && i < high {
			i++
		}
		for A[j] > pivot && j > low {
			j--
		}
		if i < j {
			// temp := A[i]
			// A[i] = A[j]
			// A[j] = temp
			A[i], A[j] = A[j], A[i]
		}
	}
	A[low] = A[j]
	A[j] = pivot

	return j
}

func quicksort(A []int, low, high int) {
	if low < high {
		var pivot = partition(A, low, high)
		quicksort(A, low, pivot-1)
		quicksort(A, pivot+1, high)
	}
}

func main() {
	A := []int{5, 9, 6, 2, 4, 3, 1}
	n := len(A)
	fmt.Println("orignal array:", A)
	quicksort(A, 0, n-1)
	fmt.Println("sorting array:", A)

}
