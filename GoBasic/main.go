package main

import "fmt"

// func main() {
// var a = [3][3]int{{2, 3, 5}, {2, 5, 6}, {2, 6, 7}}
// fmt.Println(a)
// var i, j int
// for i = 0; i < 3; i++ {
// 	for j = 0; j < 3; j++ {
// 		fmt.Print(a[i][j])
// 	}
// 	fmt.Println()
// }

// var nums = [5]int{7, 6, 5}
// target := 11
// for i, right := range nums {
// 	for j, left := range nums {
// 		if right+left == target && i != j {
// 			fmt.Println(i, j)
// 		}
// 	}
// }
// return nil

// names := []string{
// 	"Vishal",
// 	"Vinayak",
// 	"Suhas",
// 	"Nandu",
// }
// fmt.Println(names)
// sl1 := names[0:2]
// sl2 := names[1:3]
// fmt.Println(sl1, sl2)
// sl2[0] = "zzz"
// sl3 := append(sl2, "moin")
// fmt.Println(sl3)
// fmt.Println(sl1, sl2)
// fmt.Println(names)
// for i, first := range names {
// 	for j, second := range names {
// 		if first != second && second != first && i != j && j != i {
// 			fmt.Println(string(first[0:4]))
// 			_ = second
// 		}
// 	}

// }

// }
type Person struct {
	name  string
	age   int
	email string
}

func main() {
	var p1 Person
	var p2 Person

	p1.name = "Vinayak"
	p1.age = 18
	p1.email = "vinayak@gmail.com"

	p2.name = "Suhas"
	p2.age = 18
	p2.email = "suhas@gmail.com"

	fmt.Println("Name:", p1.name)
	fmt.Println("age:", p1.age)
	fmt.Println("email:", p1.email)

	fmt.Println("Name:", p2.name)
	fmt.Println("age:", p2.age)
	fmt.Println("email:", p2.email)
}
