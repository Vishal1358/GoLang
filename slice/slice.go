package main

import "fmt"

func main() {
	var myarray = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var myslice []int = myarray[1 : 5+1]
	var myslice1 []int = myslice[1 : 3+2]
	fmt.Println(myslice1, myslice)

}
