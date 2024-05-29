package main

import (
	"os"

	gotodo "github.com/WalkableGalaxy9/go-todo"
)

func main() {

	gotodo.TodoList = []gotodo.TodoItem{
		{
			Title:    "Do laundry",
			Complete: false,
		},
	}

	for true {
		gotodo.PrintTodo(os.Stdout)

		gotodo.AddTodoInput(os.Stdin, os.Stdout)
	}

}
