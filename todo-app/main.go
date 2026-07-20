package main

import (
	"bufio"
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

	res := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	fmt.Println("Welcome To TODO-APP!")
	fmt.Println("Do you want to?")
	for option != 3 {

		fmt.Println("1. View your recorded tasks and their status.")
		fmt.Println("2. Add a new task")
		fmt.Println("3. Exit")

		fmt.Scanln(&option)

		switch option {
		case 1:
			if len(todoList) == 0 {
				fmt.Println("You have nothing in your todo list, Try adding a new task!")
				

			} else {
				fmt.Fprintf(res, "ID\tTASK\tSTATUS\n")
				fmt.Fprintln(res, "--\t----\t------")

				for _, items := range todoList {

					fmt.Fprintf(res, "%d\t%s\t%t\n", items.ID, items.Task, items.Status)
				}
				res.Flush()
				fmt.Println()
			}
		case 2:
			fmt.Println("Note down a task")
			// create a scanner that listens directly to standard input
			scanner := bufio.NewScanner(os.Stdin)
			// pause the CLI program and waits for the user to type something and press Enter
			scanner.Scan()
			// save the user input in a variable
			plan = scanner.Text()
			idcount++
			todo := Todo{
				ID:     idcount,
				Task:   plan,
				Status: false,
			}
			todoList = append(todoList, todo)
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
