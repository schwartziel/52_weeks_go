package main

import(
	"fmt"
	"bufio"
    "os"
	"strings"
)

func main(){
	// there are better ways to store the to do list items, like for example, they could have due dates, and descriptions, titles etc.
	// here, it is just a list with no due date, you can remove, add things etc. There is room to improve, but the overall idea
	// of this project is just to build things as fast as possible, so here, there isn't much functionality, but perhaps in the future, that can change.
	// start with a slice that represents the items in the todo list
	var items_todo []string
	items_todo = make([]string, 3)

	keyboard := bufio.NewReader(os.Stdin)
	fmt.Println("Basic ToDo List. Enter in one thing that you need to do to start.")
	fmt.Print("Add thing to do => ")
	text, _ := keyboard.ReadString('\n')
	items_todo = append(items_todo, text)
	fmt.Println("Starting To Do List >>", items_todo)

	fmt.Println("What would you like to do next?")
	var options []string
	options = make([]string, 3)
	options[0] = "add new item"
	options[1] = "complete item"
	options[2] = "print todo list"
	fmt.Printf("%s\n%s\n%s\n", options[0], options[1], options[2])

	fmt.Print("Choose an option =>")
	text, _ = keyboard.ReadString('\n')
	text = strings.ToLower(strings.TrimSpace(text))
	// test

	if text == options[0] {
		fmt.Println("Option:", options[0])
		fmt.Print("Add thing to do => ")
		newItem, _ := keyboard.ReadString('\n')
		newItem = strings.TrimSpace(newItem)
		if newItem != "" {
			items_todo = append(items_todo, newItem)
		}
		fmt.Println("Updated To Do List >>", items_todo)
	} else if text == options[1] {
		fmt.Println("Option:", options[1])
		// naive “complete last item” keeping your simple flow
		if len(items_todo) > 0 {
			items_todo = items_todo[:len(items_todo)-1]
		}
		fmt.Println("Updated To Do List >>", items_todo)
	} else if text == options[2] {
		fmt.Println("Option:", options[2])
		fmt.Println("To Do List >>", items_todo)
	} else {
		fmt.Println("Unknown option.")
	}

}
