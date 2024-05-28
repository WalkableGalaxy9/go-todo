package main

import (
	"os"

	gotodo "github.com/WalkableGalaxy9/go-todo"
)

func main() {

	todo := gotodo.TodoItem{
		Title:    "Do laundry",
		Complete: false,
	}

	gotodo.PrintTodo(os.Stdout, todo)
}
