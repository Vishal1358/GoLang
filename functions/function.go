package main

import (
	"fmt"
	"math"
)

func f1() {
	fmt.Println("This is f1() function.")
}

func f2(a int, b int) {
	fmt.Println("Sum:", a+b)
}
func f3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}

func f4(a float64) float64 {
	return math.Pow(a, a)
}

func main() {
	f1()
	f2(5, 6)
	f3(1, 2, 3, 4.5, 4.6, "Vishal")
	p := f4(5.1)
	fmt.Printf("%.2f", p)
}
