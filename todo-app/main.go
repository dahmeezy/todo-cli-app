package main

import (
	"bufio"
	"fmt"
	"os"
)

type Todo struct {
	ID     int
	Task   string
	Status bool
}

func main() {

	idcount := 0

	usage := "Usage: go run . todo-app"
	fmt.Println(len(os.Args))
	fmt.Println(os.Args)

	if len(os.Args) < 2 || len(os.Args) > 2 {
		fmt.Println(usage)
		return
	}

	if os.Args[1] != "todo-app" {
		fmt.Println(usage)
		return
	}

	var todo int
	var plan string
	var todoList []Todo

	fmt.Println("Welcome To TODO-APP!")
	fmt.Println("Do you want to?")
	fmt.Println("1. View your recorded tasks and their status.")
	fmt.Println("2. Add a new task")
	fmt.Println("3. Exit")

	fmt.Scanln(&todo)

	switch todo {
	case 1:
		if len(todoList) == 0 {
			fmt.Println("You have nothing in your todo list,")
			fmt.Println("2. add a new list or,")
			fmt.Println("3. Exit?")
		}
	case 2:
		fmt.Println("Note down a task")
		// create a scanner that listens directly to standard input
		scanner := bufio.NewScanner(os.Stdin)
		// pause the CLI program and waits for the user to type something and press Enter
		scanner.Scan()
		// save the user input in a variable
		plan += scanner.Text()
		idcount++
		todos := Todo{
			ID:     idcount,
			Task:   plan,
			Status: false,
		}
		todoList = append(todoList, todos)
		fmt.Println("Add another task?")
		fmt.Scanln(&plan)


	case 3:
		return
	default:
		fmt.Println("Choose a valid option")
	}

}
