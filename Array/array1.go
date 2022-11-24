package main

import "fmt"

func main() {
	var start, end int
	fmt.Println("Start value: ")
	fmt.Scan(&start)
	fmt.Println("End value: ")
	fmt.Scan(&end)
	var a int = 0
	for i := start; i < end+1; i++ {
		for j := i; j < i+1; j++ {
			if i%j == 0 {
				a = a + 1
			}
		}
		if a == 2 {
			fmt.Println(i)
		}
	}
}
