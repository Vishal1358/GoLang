package main

import "fmt"

func main() {

	fmt.Println("Enter age for voting : ")
	var age int
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("You are able to voting")
	} else {
		fmt.Println("You are not able to voting")
	}

}
