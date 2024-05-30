package gotodo

import "fmt"

type TodoItem struct {
	Title    string
	Complete bool
}

var TodoList []TodoItem

func DeleteTodo(indexToRemove int) error {

	if indexToRemove > len(TodoList) {
		return fmt.Errorf("index out of range")
	} else {
		TodoList = append(TodoList[:indexToRemove-1], TodoList[indexToRemove:]...)
	}

	return nil
}

func CreateTodo(title string) {
	item := TodoItem{Title: title, Complete: false}

	TodoList = append(TodoList, item)
}

func ToggleTodo(index int) error {

	TodoList[index-1].Complete = !TodoList[index-1].Complete

	return nil
}
