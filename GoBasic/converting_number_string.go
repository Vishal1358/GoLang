package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := string(99)
	fmt.Println(s)
	// converting 99 to c means ascii value

	//s1 := strring(44.2)
	var myStr = fmt.Sprintf("%f", 44.2)
	fmt.Println(myStr)

	var myStr1 = fmt.Sprintf("%d", 34234)
	fmt.Println(myStr1)
	//it will representing the int value to string

	fmt.Println(string(34234))
	//it will convert int value string uniqcode

	s1 := "3.123"
	fmt.Printf("%T\n", s1)

	//now convert string to int/float
	// We can use here strconv package and ParsrFloat function
	var f1, err = strconv.ParseFloat(s1, 64)
	_ = err
	fmt.Println(f1)
	fmt.Println(f1 * 3.4)

	//We can do here Ascii to int and int to ascii values
	// Atoi ===> Ascii to int
	i, err := strconv.Atoi("-50")
	// Itoa ===> int to ascii
	s2 := strconv.Itoa(20)

	fmt.Printf("i type is %T, i value is %v\n", i, i)
	fmt.Printf("s2 type is %T, s2 value is %v\n", s2, s2)

	var x = fmt.Sprintf("%d", 34234)
	fmt.Println("%T", x)

}
