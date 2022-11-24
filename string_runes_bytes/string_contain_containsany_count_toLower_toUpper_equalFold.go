package main

import (
	"fmt"
	"strings"
)

func main() {
	p := fmt.Println

	result := strings.Contains("Hi goodevening vishal", "vishal")
	p(result)

	result = strings.ContainsAny("success", "xys")
	p(result)

	n := strings.Count("Cheese", "e")
	p(n)

	n1 := strings.Count("fives", "")
	p(n1)

	p(strings.ToLower("HI VINAYAK"))
	p(strings.ToUpper("HI vinayak"))

	p("GO" == "GO")
	p("go" == "GO")
	p(strings.ToLower("Go") == strings.ToLower("go"))

	p(strings.EqualFold("GO", "go"))

	myStr := strings.Repeat("ab", 10)
	p(myStr)

	myStr = strings.Replace("192.168.0.1.2", ".", ":", 2)
	p(myStr)

	// myStr = strings.Replace("192.168.0.1.2", ".", ":", -1)
	// p(myStr)

	myStr = strings.ReplaceAll("192.168.0.1", ".", ":")
	p(myStr)

	s := strings.Split("a,b,c", ",")
	fmt.Printf("%T\n", s)
	fmt.Printf("%#v\n", s)

	s = strings.Split("Go for Go!", "")
	fmt.Printf("%#v\n", s)

	s = []string{"I", "learn", "Golang"}
	myStr = strings.Join(s, "-")
	p(myStr)

	s1 := strings.TrimSpace("\t Goodby Windows, Welcome Linux! \n")
	fmt.Printf("%q\n", s1)

	s2 := strings.Trim("...Hello Gophers!!!!?", ".!?")
	fmt.Printf("%q\n", s2)

}
