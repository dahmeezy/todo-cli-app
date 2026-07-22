package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

const fileName = "todos.json"

type DATA struct {
	ID     int    `json:"id"`
	Task   string `json:"task"`
	Status string `json:"status"`
}
type Todo struct {
	ID     int
	Task   string
	Status string
}

func main() {

	idcount := 0

	usage := "Usage: go run . todo-app"

	if len(os.Args) < 2 || len(os.Args) > 2 {
		fmt.Println(usage)
		return
	}

	if os.Args[1] != "todo-app" {
		fmt.Println(usage)
		return
	}

	var option int
	var todoList []Todo
	var comp int
	var del int
	result := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	fmt.Println("Welcome To TODO-APP!")
	fmt.Println("Do you want to?")
	for option != 5 {

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
			todoList = addNewTask(idcount, todoList)
			fmt.Println("Your task has been added!")
			fmt.Println()

		case 3:
			fmt.Println("Enter Task ID of completed task:")
			fmt.Scanln(&comp)

			markTaskAsCompleted(comp, todoList)
			fmt.Println()
			printMenu()

		case 4:
			fmt.Println("Enter task ID of task to be deleted:")
			fmt.Scanln(&del)
			todoList = deleteATask(del, todoList)

		case 5:
			return
		default:
			fmt.Println("Choose a valid option!")
			fmt.Println()
		}
	}

}
