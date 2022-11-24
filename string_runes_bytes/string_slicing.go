package main

import "fmt"

func main() {
	s1 := "Hi i'am vishal"
	fmt.Println(s1[8:14])

	rs := []rune(s1)
	fmt.Printf("%T\n", rs)
}
