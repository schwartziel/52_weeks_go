package main

import(
	"fmt"
	"rsc.io/quote"
)

func main() {
fmt.Println("hello, world.")
	fmt.Println("From the go documentation on getting started. Printing out a quote:")
	fmt.Println(quote.Go())

	fmt.Print("\n")
	fmt.Println("Working with numbers and operators...")
	fmt.Print("\n")
	// Hello reader, before we begin, this is a single line comment.
	/*
    this is
	a
	multiple
	line
	comment
	*/
	// Just like C.
	fmt.Println("working with numbers is just how you would expect. you can use +, -, *, / to specify algerbraic equations.")
	number_one := 14
	number_two := 3
	number_3 := number_one + number_two
	number_7 := number_one - number_two
	number_4 := number_one * number_two
	number_5 := number_one / number_two
	number_6 := number_one % number_two
	fmt.Println("number_one is:", number_one)
	fmt.Println("number_two is:", number_two)
	fmt.Println("the first plus the seccond number is: ", number_3)
	fmt.Println("the first minus the seccond number is: ", number_7)
	fmt.Println("the first times the seccond number is: ", number_4)
	fmt.Println("the first divided by the seccond number is: ", number_5)
	fmt.Println("the remainder of the first number divided by the seccond is: ", number_6)

	fmt.Print("\n")
	fmt.Println("Representing integers in go...")
	my_number := 42
	fmt.Println(" using my_number := 42 ", my_number)
	my_number = 4_2
	fmt.Println(" using my_number := 4_2", my_number)
	my_number = 0600
	fmt.Println(" using my_number := 0600 ", my_number)
	my_number = 0b1010
	fmt.Println(" using my_number := 0b1010 ", my_number)
	my_number = 0xBadFace
	fmt.Println(" using my_number := 0xBadFace ", my_number)
	fmt.Println("Giving up on this program because I got bored")
}
