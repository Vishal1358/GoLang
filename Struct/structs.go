package main

import "fmt"

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	pl("Structs in golang")

	vishal := User{"Vishal", "vishal@gmail.com", true, 26}
	pl(vishal)
	pf("Vishal details are %+v : \n", vishal)
	pf("Name is %v Email id is %v Age is %v.", vishal.Name, vishal.Email, vishal.Age)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
