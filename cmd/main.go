package main

import (
	"os"

	gotodo "github.com/WalkableGalaxy9/go-todo"
)

func main() {

	todo := []gotodo.TodoItem{
		{
			Title:    "1. Do laundry",
			Complete: false,
		},
		{
			Title:    "2. Learn go",
			Complete: false,
		},
		{
			Title:    "3. Go to dentist",
			Complete: false,
		},
		{
			Title:    "4. Update Driving licence",
			Complete: true,
		},
		{
			Title:    "5. Make seating plan",
			Complete: false,
		},
		{
			Title:    "6. Send remaining invites",
			Complete: false,
		},
		{
			Title:    "7. Plan political rally",
			Complete: true,
		},
		{
			Title:    "8. Another item",
			Complete: false,
		},
		{
			Title:    "9. Clean kitchen",
			Complete: false,
		},
		{
			Title:    "10. Blah",
			Complete: false,
		},
	}

	gotodo.PrintTodo(os.Stdout, todo)
}
