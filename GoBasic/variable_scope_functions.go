package main

import "fmt"

var a = 3

func demo() {
	// var a = 2
	var greet = "Good Morning :)"
	fmt.Printf("%v Vishal", greet)
	// fmt.Println("demo func", a)
	// the function allways preferences give the own variable or local variable

}

func main() {
	demo()
	// print("main func", a)
	// if the variable is not avalable inside the functions then it will go the global or package variable
}
