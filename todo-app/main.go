package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

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
	var todoList string

	fmt.Println("Welcome To TODO-APP!")
	fmt.Println("Do you want to?")
	fmt.Println("1. View your recorded tasks and their status.")
	fmt.Println("2. Add a new task")
	fmt.Println("3. Exit")

	fmt.Scanln(&todo)

	switch todo {
	case 1:
		if todoList == "" {
			fmt.Println("You have nothing in your todo list,")
			fmt.Println("2. add a new list or,")
			fmt.Println("3. Exit?")
		}
	case 2:
		fmt.Println("Note down a task")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		plan += scanner.Text()
	case 3:
		return
	default:
		fmt.Println("Choose a valid option")
	}
}
