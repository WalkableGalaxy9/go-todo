package gotodo

import (
	"encoding/json"
	"fmt"
	"io"
)

type TodoItem struct {
	Title    string
	Complete bool
}

func PrintTodo(writer io.Writer, items []TodoItem) {

	for index, item := range items {
		completeString := "complete"

		if !item.Complete {
			completeString = "incomplete"
		}

		fmt.Fprintf(writer, "%d. %s - %s\n", index+1, item.Title, completeString)
	}
}

func CreateList(items ...string) (todos []TodoItem) {

	for _, item := range items {
		todos = append(todos, TodoItem{item, false})
	}

	return

}

func PrintTodoJSON(writer io.Writer, items []TodoItem) {

	json.NewEncoder(writer).Encode(items)
}
