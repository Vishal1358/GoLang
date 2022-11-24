package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m)
	m["Key1"] = 10
	m["Key2"] = 20
	m["Key3"] = 30
	fmt.Println(m)
	m["Key2"] = 555
	m["Key4"] = 555
	fmt.Println(m)
}
