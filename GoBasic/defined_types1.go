package main

import "fmt"

type km float64

type mile float64

func main() {

	type speed uint

	var s1 speed = 10
	var s2 speed = 20

	fmt.Println(s2 - s1)
	//because it is same type

	var x uint
	//x = s1 // it will give compile error defferent types

	x = uint(s1)
	_ = x

	var s3 speed = speed(x)
	_ = s3

	var parisToLondon km = 465
	var distanceInMiles mile

	distanceInMiles = mile(parisToLondon / 0.62)
	fmt.Println(distanceInMiles)
}
