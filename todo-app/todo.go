package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"text/tabwriter"
)

func loadTodo() ([]Todo, error) {
	data, er := os.ReadFile("todos.json")
	if er != nil {
		if !errors.Is(er, os.ErrNotExist) {
			return nil, er
		} else {
			return []Todo{}, nil
		}
	}
	var todos []Todo
	err := json.Unmarshal(data, &todos)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

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
		Status: "PENDING",
	}
	todoList = append(todoList, todo)
	return todoList
}

func viewTask(todoList []Todo, result *tabwriter.Writer) {
	fmt.Fprintf(result, "ID\tTASK\tSTATUS\n")
	fmt.Fprintln(result, "--\t----\t------")

	for _, items := range todoList {

		fmt.Fprintf(result, "%d\t%s\t%s\n", items.ID, items.Task, items.Status)
	}
	result.Flush()
}

func printMenu() {
	fmt.Println("1. View your recorded tasks and their status.")
	fmt.Println("2. Add a new task")
	fmt.Println("3. Mark task as completed")
	fmt.Println("4. Delete a task")
	fmt.Println("5. Exit")
}

func markTaskAsCompleted(comp int, todoList []Todo) {
	exists := false
	for i := range todoList {
		item := &todoList[i]
		if item.ID == comp {
			exists = true
			item.Status = "COMPLETED"
			break
		}
	}
	if !exists {
		fmt.Println("There's no task with that ID!")
	} else {
		fmt.Println("Task status updated. Do you want to:")
	}
}

func deleteATask(del int, todoList []Todo) []Todo {
	exists := false
	var res string
	for i := range todoList {
		item := &todoList[i]
		if item.ID == del {
			todoList = append(todoList[:i], todoList[i+1:]...)
			exists = true
			break
		}
	}

	if !exists {
		res = "There's no task with that ID!"
	} else {
		res = fmt.Sprintf("Task %d deleted. Do you want to:", del)
	}

	fmt.Println(res)
	return todoList
}

func saveTodo() error {
	var todos []Todo
	data, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0777)
}
