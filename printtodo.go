package gotodo

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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

func WriteJSONToFile(filename string, items []TodoItem) error {

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	PrintTodoJSON(file, items)

	return nil
}
