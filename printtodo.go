package gotodo

import (
	"fmt"
	"io"
)

type TodoItem struct {
	Title    string
	Complete bool
}

func PrintTodo(writer io.Writer, items []TodoItem) {

	for _, item := range items {
		completeString := "complete"

		if !item.Complete {
			completeString = "incomplete"
		}

		fmt.Fprintf(writer, "%s - %s\n", item.Title, completeString)
	}
}

func CreateList(items ...string) (todos []TodoItem) {

	for _, item := range items {
		todos = append(todos, TodoItem{item, false})
	}

	return

}
