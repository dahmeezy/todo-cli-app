package main

import (
	"fmt"
	"os"
	"text/tabwriter"
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

	var option int
	var plan string
	var todoList []Todo

	result := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	fmt.Println("Welcome To TODO-APP!")
	fmt.Println("Do you want to?")
	for option != 3 {

		printMenu()

		fmt.Scanln(&option)

		switch option {
		case 1:
			if len(todoList) == 0 {
				fmt.Println("You have nothing in your todo list, Try adding a new task!")

			} else {
				viewTask(todoList, result)
				fmt.Println()
			}
		case 2:
			fmt.Println("Note down a task")
			idcount++
			todoList = addNewTask(idcount, plan, todoList)
			fmt.Println("Your task has been added!")
			fmt.Println()
		case 3:
			return
		default:
			fmt.Println("Choose a valid option!")
			fmt.Println()
		}
	}

}
