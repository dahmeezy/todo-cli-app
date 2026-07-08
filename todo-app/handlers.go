package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Id   int
	Task string
	Done bool
}

type Data struct {
	Todos []Todo
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/todo.html")
	if err != nil {
		http.Error(w, "500 - Internal Server Error!", http.StatusInternalServerError)
		return
	}

	todos, er := getAllTodos()
	if er!=nil{
		return
	}
	tmpl.Execute(w, Data{Todos: todos})
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "400 -Bad Request", http.StatusBadRequest)
		return
	}

	task := r.FormValue("task")

	e := addTodo(task)
	if e != nil {
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
