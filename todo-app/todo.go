package main

import (
	"bufio"
	"fmt"
	"os"
	"text/tabwriter"
)

func addNewTask(idcount int, todoList []Todo) []Todo {
	// create a scanner that listens directly to standard input
	scanner := bufio.NewScanner(os.Stdin)
	// pause the CLI program and waits for the user to type something and press Enter
	scanner.Scan()
	// save the user input in a variable
	plan := scanner.Text()
	// idcount++
	todo := Todo{
		ID:     idcount,
		Task:   plan,
		Status: false,
	}
	todoList = append(todoList, todo)
	return todoList
}

func viewTask(todoList []Todo, result *tabwriter.Writer) {
	fmt.Fprintf(result, "ID\tTASK\tSTATUS\n")
	fmt.Fprintln(result, "--\t----\t------")

	for _, items := range todoList {

		fmt.Fprintf(result, "%d\t%s\t%t\n", items.ID, items.Task, items.Status)
	}
	result.Flush()
}

func printMenu() {
	fmt.Println("1. View your recorded tasks and their status.")
	fmt.Println("2. Add a new task")
	fmt.Println("3. Exit")
}
