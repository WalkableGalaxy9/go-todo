package main

import (
	"fmt"
	"io"
	"os"

	gotodo "github.com/WalkableGalaxy9/go-todo"
)

func main() {

	todoCLIViewModel := gotodo.TodoViewModelCLI{}
	todoCLIViewModel.TodoList = gotodo.TodoList{
		{
			Title:    "Do laundry",
			Complete: false,
		},
	}

	for {
		ManageTodoList(&todoCLIViewModel, os.Stdin, os.Stdout)
	}

}

func ManageTodoList(viewModel gotodo.TodoViewModel, input io.Reader, output io.Writer) {
	fmt.Fprint(output, "\033[H\033[2J")

	viewModel.PrintTodo(output)

	option := gotodo.GetMenuOption(input, output)

	switch option {
	case gotodo.MenuAddTodo:
		viewModel.AddTodoInput(input, output)
	case gotodo.MenuDelete:
		viewModel.DeleteTodoInput(input, output)
	case gotodo.MenuToggle:
		viewModel.ToggleTodoInput(input, output)
	case gotodo.MenuExit:
		os.Exit(0)
	}
}
