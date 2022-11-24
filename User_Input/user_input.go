package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// var myString string
	// fmt.Println("Enter Your Name ...?")
	// fmt.Scanln(&myString)
	// fmt.Println("Employee name: ", myString)

	// var name string = "Vishal"
	// var a, b = fmt.Println(name)
	// fmt.Println(a, b)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter you full name: ")
	// myName, _ := reader.ReadString('\n')
	// fmt.Println(myName)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your rating: ")
	myRating, _ := reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(myRating), 64)
	fmt.Println(mynumRating + 2)
}
