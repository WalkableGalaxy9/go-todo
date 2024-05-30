package gotodo

import (
	"fmt"
)

type TodoItem struct {
	Title    string
	Complete bool
}

type TodoList []TodoItem

func (t *TodoList) DeleteTodo(indexToRemove int) error {

	if indexToRemove > len(*t) {
		return fmt.Errorf("index out of range")
	} else {
		*t = append((*t)[:indexToRemove-1], (*t)[indexToRemove:]...)
	}

	return nil
}

func (t *TodoList) CreateTodo(title string) {
	item := TodoItem{Title: title, Complete: false}

	*t = append(*t, item)
}

func (t *TodoList) ToggleTodo(index int) error {

	if index > len(*t) {
		return fmt.Errorf("index out of range")
	} else {
		(*t)[index-1].Complete = !(*t)[index-1].Complete
	}
	return nil
}
