package main

import "fmt"

func main() {
	var employees map[string]string
	fmt.Printf("%#v\n", employees)

	fmt.Printf("No of pairs:%d\n", len(employees))

	fmt.Printf("The value for key %q is %q", "Dan", employees["Dan"])

	var accounts map[string]float64
	fmt.Printf("%#v\n", accounts["0x323"])

	var m1 map[[5]int]string
	_ = m1

	// employees["Dan"] = "Programmer"

	people := map[string]float64{}

	people["John"] = 21.4
	people["Marry"] = 10
	fmt.Println(people)

	map1 := make(map[int]int)
	map1[4] = 8

	balances := map[string]float64{
		"USD": 323.11,
		"EUR": 432.13,
		// 50:  322.1,  //error,all keys the same type
		"CHF": 3243.1,
	}
	fmt.Println(balances)

}
