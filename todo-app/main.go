package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

const fileName = "todos.json"

type Todo struct {
	ID     int    `json:"id"`
	Task   string `json:"task"`
	Status string `json:"status"`
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
	printMenu()
	for option != 5 {

		fmt.Scanln(&option)

		switch option {
		case 1:
			if len(todoList) == 0 {
				fmt.Println("You have nothing in your todo list, Try adding a new task!")
				printMenu()
			} else {
				viewTask(todoList, result)
				fmt.Println()
				printMenu()
			}
		case 2:
			fmt.Println("Note down a task")
			idcount++
			todoList = addNewTask(idcount, todoList)
			fmt.Println("Your task has been added!")
			fmt.Println()
			saveTodo()
			printMenu()

		case 3:
			fmt.Println("Enter Task ID of completed task:")
			fmt.Scanln(&comp)

			markTaskAsCompleted(comp, todoList)
			fmt.Println()
			saveTodo()
			printMenu()

		case 4:
			fmt.Println("Enter task ID of task to be deleted:")
			fmt.Scanln(&del)
			todoList = deleteATask(del, todoList)
			fmt.Println()
			saveTodo()
			printMenu()

		case 5:
			saveTodo()
			return
		default:
			fmt.Println("Choose a valid option!")
			printMenu()
			fmt.Println()
		}
	}

}
