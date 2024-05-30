package main

import (
	"fmt"
	"io"
	"os"

	gotodo "github.com/WalkableGalaxy9/go-todo"
)

func main() {

	todovm := gotodo.TodoViewModelCLI{}
	todovm.TodoList = gotodo.TodoList{
		{
			Title:    "Do laundry",
			Complete: false,
		},
	}

	for {
		ManageTodoList(&todovm, os.Stdin, os.Stdout)
	}

}

func ManageTodoList(vm gotodo.TodoViewModel, input io.Reader, output io.Writer) {
	fmt.Fprint(output, "\033[H\033[2J")

	vm.PrintTodo(output)

	option := gotodo.GetMenuOption(input, output)

	switch option {
	case gotodo.MenuAddTodo:
		vm.AddTodoInput(input, output)
	case gotodo.MenuDelete:
		vm.DeleteTodoInput(input, output)
	case gotodo.MenuToggle:
		vm.ToggleTodoInput(input, output)
	case gotodo.MenuExit:
		os.Exit(0)
	}
}
