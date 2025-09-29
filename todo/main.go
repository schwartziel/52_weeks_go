package main

import(
	"fmt"
	"os"
)

func main(){
	fmt.Println("Yes, this is yet another to do app written in go, there are many like it, but I did not use a tutorial to build this. Just my own raw brain power.")
	fmt.Println("Yes, that also means it's likely not going to work very well.")
	// we're going to choose to store our list in a "slice" in go. So I have heard on stack overflow, anytime we think of putting something in a list, we should use a slice.
	// init the slice.
	// i changed my mind let's make it an object of slices.
	var todo_items []string // creates a string variable
	todo_items = make([]string, 5) // sets it to a slice with length and capacity 5
	var input string
	type todo_list string; { // struct
		var item_titles string
	    var item_description string
	    var item_duedate string
}
	fmt.Printf("%s\nSelect an action to take from below:\n")
	fmt.Printf("%s  add (add an item to the list)\n")
	fmt.Printf("%s  remove (remove an item from the list.) \n ")
	fmt.Printf("%s  check (mark an item as complete.)\n")
	fmt.Print("selector => ")
	fmt.Scanln(&input)
	fmt.Printf("%s", input)
	if input == "add" {
		fmt.Printf("%s adding item.\n")
		fmt.Printf("%s\ntitle: ")
		fmt.Scanln(&todo_list.item_titles)
		fmt.Printf("%s\ndescription: ")
		var item_description string
		fmt.Scanln(&item_description)
		fmt.Printf("%s\ndue date:")
		var item_duedate string
		fmt.Scanln(&item_duedate)

	}
}
