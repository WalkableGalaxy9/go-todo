package main

import (
	"fmt"
	"os"

	gotodo "github.com/WalkableGalaxy9/go-todo"
)

func main() {

	todos := gotodo.TodoList{
		{
			Title:    "Do laundry",
			Complete: false,
		},
	}

	for {
		fmt.Fprint(os.Stdout, "\033[H\033[2J")
		gotodo.PrintTodo(&todos, os.Stdout)

		option := gotodo.GetMenuOption(os.Stdin, os.Stdout)

		switch option {
		case gotodo.MenuAddTodo:
			gotodo.AddTodoInput(&todos, os.Stdin, os.Stdout)
		case gotodo.MenuDelete:
			gotodo.DeleteTodoInput(&todos, os.Stdin, os.Stdout)
		case gotodo.MenuToggle:
			gotodo.ToggleTodoInput(&todos, os.Stdin, os.Stdout)
		case gotodo.MenuExit:
			os.Exit(0)
		}
	}

}
