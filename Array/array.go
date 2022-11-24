package main

import "fmt"

func main() {
	// a := [5]int{1, 2, 3, 4, 5}	// single Dimensional Arrays
	var a = [2][2]int{{1, 2}, {2, 5}}       // single Dimensional Arrays
	var b = [3][3]int{{1, 2, 4}, {2, 5, 9}} // single Dimensional Arrays
	var c = [2][4]int{{1, 2}, {2, 5, 8, 9}} // single Dimensional Arrays
	fmt.Printf("single Dimensional Arrays %v\n,multi Dimensional Arrays %v\n, multi Dimensional Arrays %v\n", a, b, c)

}
