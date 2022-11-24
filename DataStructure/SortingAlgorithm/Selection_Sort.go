package main

import "fmt"

func SelectionSort(items []int) {
	for i := 0; i < len(items)-1; i++ {
		min := i
		for j := i + 1; j < len(items); j++ {
			if items[min] > items[j] {
				min = j
			}
		}
		temp := items[i]
		items[i] = items[min]
		items[min] = temp
	}
}

func main() {
	items := []int{25, -2, 660, 8, -1}
	fmt.Println(items)
	SelectionSort(items)
	fmt.Println(items)

}
