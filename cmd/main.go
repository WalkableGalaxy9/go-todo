package main

import (
	"fmt"
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
		fmt.Fprint(os.Stdout, "\033[H\033[2J")
		gotodo.PrintTodo(os.Stdout)

		option := gotodo.GetMenuOption(os.Stdin, os.Stdout)

		switch option {
		case gotodo.MenuAddTodo:
			gotodo.AddTodoInput(os.Stdin, os.Stdout)
		case gotodo.MenuDelete:
			gotodo.DeleteTodoInput(os.Stdin, os.Stdout)

		}

	}

}
