package gotodo

import (
	"fmt"
	"io"
)

type TodoItem struct {
	Title    string
	Complete bool
}

func PrintTodo(writer io.Writer, item TodoItem) {

	completeString := "complete"

	if !item.Complete {
		completeString = "incomplete"
	}

	fmt.Fprintf(writer, "%s - %s", item.Title, completeString)
}
