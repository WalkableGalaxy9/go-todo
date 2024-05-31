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

func (t *TodoList) CreateTodo(newTodoTitle string) {
	newTodo := TodoItem{Title: newTodoTitle, Complete: false}

	*t = append(*t, newTodo)
}

func (t *TodoList) ToggleTodo(indexToToggle int) error {

	if indexToToggle > len(*t) {
		return fmt.Errorf("index out of range")
	} else {
		(*t)[indexToToggle-1].Complete = !(*t)[indexToToggle-1].Complete
	}
	return nil
}
