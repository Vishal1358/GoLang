package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Sturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case today + 3:
		fmt.Println("In three days.")
	default:
		fmt.Println("Too far away.")
	}
	fmt.Println(time.Now())
}
