package main

import (
	"fmt"
)

func main() {
	var ival int = 10
	var fval float32 = 23.34
	var cval complex64 = 1 + 2i

	bytes, err := fmt.Printf("%d %g %g\n", ival, fval, cval)
	fmt.Println(bytes, err)
}
