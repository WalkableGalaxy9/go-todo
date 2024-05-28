package main

import (
	"os"

	gotodo "github.com/WalkableGalaxy9/go-todo"
)

func main() {

	todo := []gotodo.TodoItem{
		{
			Title:    "Do laundry",
			Complete: false,
		},
		{
			Title:    "Learn go",
			Complete: false,
		},
		{
			Title:    "Go to dentist",
			Complete: false,
		},
		{
			Title:    "Update Driving licence",
			Complete: true,
		},
		{
			Title:    "Make seating plan",
			Complete: false,
		},
		{
			Title:    "Send remaining invites",
			Complete: false,
		},
		{
			Title:    "Plan political rally",
			Complete: true,
		},
		{
			Title:    "Another item",
			Complete: false,
		},
		{
			Title:    "Clean kitchen",
			Complete: false,
		},
		{
			Title:    "Blah",
			Complete: false,
		},
	}

	gotodo.PrintTodo(os.Stdout, todo)

	gotodo.PrintTodoJSON(os.Stdout, todo)

	gotodo.WriteJSONToFile("../output/todo.txt", todo)

	gotodo.ReadJSONFromAFile(os.DirFS("../input"), "todo.txt", os.Stdout)
}
