package main

import "fmt"

func f1(s []int) {
	s[0] = 100

}

func main() {
	s := []int{1, 2}
	f1(s)
	fmt.Println(s)
}
