package main

import "fmt"

func main() {
	type book struct {
		title, author string
		year          int
	}

	lastbook := book{title: "Why I Killed Gandhi"}
	fmt.Println(lastbook.title)

	fmt.Printf("%#v\n", lastbook)

	lastbook.author = "Nathuram Godse"
	lastbook.year = 2022
	fmt.Printf("%+v\n", lastbook)

	aBook := book{title: "Why I Killed Gandhi", author: "Nathuram Godse", year: 2022}
	fmt.Println(aBook == lastbook)

	randomBook := book{title: "Random Title", author: "John Doe", year: 100}
	fmt.Println(randomBook == lastbook) // => false

	// = creates a copy of a struct
	myBook := randomBook
	myBook.year = 2020              // modifying only myBook
	fmt.Println(myBook, randomBook) // => {Random Title John Doe 2020} {Random Title John Doe 100}

	vinayak := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Vinayak",
		lastName:  "Havannavar",
		age:       18,
	}

	fmt.Printf("%#v\n", vinayak)
	fmt.Printf("Vinayak's Age: %d\n", vinayak.age)

	type Book struct {
		string
		float64
		bool
	}

	b1 := Book{"The men who killed Gandhi", 200.2, false}

	fmt.Printf("%#v\n", b1)

	fmt.Println(b1.string)

	type Employee struct {
		name   string
		salary int
		bool
	}

	e := Employee{"Nandu", 100000, false}
	fmt.Printf("%#v\n", e)
	e.bool = true
	fmt.Printf("%#v\n", e)

}
