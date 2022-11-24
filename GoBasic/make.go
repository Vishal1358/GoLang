package main

import "fmt"

var pl = fmt.Println

func main() {
	a := make([]int, 6, 6)
	pl(a)
	pl("Length : %v", len(a))
	pl("Capacity : %v", cap(a))

}
