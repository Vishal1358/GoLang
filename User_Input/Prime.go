// prime number

package main

import "fmt"

func main() {
	var n, a int
	fmt.Println("Enter a number : ")
	fmt.Scan(&n)
	for i := 2; i < n+1; i++ {
		if n%i == 0 {
			a++
		}
	}
	if a == 1 {
		fmt.Println("Its a prime no")
	} else {
		fmt.Println("Its not a prime no")
	}

}
