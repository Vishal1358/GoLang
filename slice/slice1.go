package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// g := []string{"Vishal", "Vishal", "Vishal", "Vishal", "Vishal", "Vishal", "Vishal", "Vishal", "Vishal", "Vishal"}
	b := a[:]   // slice of all elements
	c := a[3:]  // slice from 4th element to end
	d := a[:6]  // slice first 6 elements
	e := a[3:6] // slice slice 4th, 5th, 6th elements
	// b := g[:]   // slice of all elements
	// c := g[3:]  // slice from 4th element to end
	// d := g[:6]  // slice first 6 elements
	// e := g[3:6] // slice slice 4th, 5th, 6th elements
	fmt.Println(a)
	fmt.Println("Length: %v\n", len(a))
	fmt.Println("Capacitty: %v\n", cap(a))
	// fmt.Println(g)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
