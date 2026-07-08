package main

import (
	"database/sql"
)

var db *sql.DB

func getAllTodos() ([]Todo, error) {
	rows, err := db.Query("SELECT id, task, done FROM todos")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var t Todo

		err := rows.Scan(&t.Id, &t.Task, &t.Done)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

func addTodo(task string) error {
	_, err := db.Exec("INSERT INTO todos (task,done) VALUES(?,?)", task, 0)
	if err != nil {
		return err
	}
	return nil
}
