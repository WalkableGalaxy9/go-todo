package gotodo

import (
	"reflect"
	"testing"
)

var Todos TodoList

func TestCreateTodo(t *testing.T) {

	want := TodoList{
		{
			"Do laundry",
			false,
		},
		{
			"Learn Go",
			false,
		},
	}

	todoList := TodoList{
		{
			"Do laundry",
			false,
		},
	}

	todoList.CreateTodo("Learn Go")

	if !reflect.DeepEqual(want, todoList) {
		t.Errorf("Wanted %v got %v", want, todoList)
	}

}

func TestDeleteTodo(t *testing.T) {

	t.Run("First item", func(t *testing.T) {
		todoList := TodoList{
			{"Do laundry", false},
			{"Go shopping", false},
			{"learn go", false},
		}

		err := todoList.DeleteTodo(1)

		want := TodoList{
			{"Go shopping", false},
			{"learn go", false},
		}

		if !reflect.DeepEqual(todoList, want) {
			t.Errorf("Got %v want %v", todoList, want)
		}

		if err != nil {
			t.Errorf("Got an error %d", err)
		}
	})

	t.Run("Second item", func(t *testing.T) {
		todoList := TodoList{
			{"Do laundry", false},
			{"Go shopping", false},
			{"learn go", false},
		}

		err := todoList.DeleteTodo(2)

		want := TodoList{
			{"Do laundry", false},
			{"learn go", false},
		}

		if !reflect.DeepEqual(todoList, want) {
			t.Errorf("Got %v want %v", todoList, want)
		}

		if err != nil {
			t.Errorf("Got an error %d", err)
		}
	})

	t.Run("Third item", func(t *testing.T) {
		todoList := TodoList{
			{"Do laundry", false},
			{"Go shopping", false},
			{"learn go", false},
		}

		err := todoList.DeleteTodo(3)

		want := TodoList{
			{"Do laundry", false},
			{"Go shopping", false},
		}

		if !reflect.DeepEqual(todoList, want) {
			t.Errorf("Got %v want %v", todoList, want)
		}

		if err != nil {
			t.Errorf("Got an error %d", err)
		}
	})

	t.Run("Index out of range", func(t *testing.T) {
		todoList := TodoList{
			{"Do laundry", false},
			{"Go shopping", false},
			{"learn go", false},
		}

		err := todoList.DeleteTodo(4)

		if err == nil {
			t.Errorf("Expected an error but didn't get one")
		}
	})
}

func TestToggleTodo(t *testing.T) {

	t.Run("First item", func(t *testing.T) {
		todoList := TodoList{
			{"Do laundry", false},
			{"Go shopping", false},
			{"learn go", false},
		}

		err := todoList.ToggleTodo(1)

		want := TodoList{
			{"Do laundry", true},
			{"Go shopping", false},
			{"learn go", false},
		}

		if !reflect.DeepEqual(todoList, want) {
			t.Errorf("Got %v want %v", todoList, want)
		}

		if err != nil {
			t.Errorf("Got an error %d", err)
		}
	})

	t.Run("Second item", func(t *testing.T) {
		todoList := TodoList{
			{"Do laundry", true},
			{"Go shopping", true},
			{"learn go", true},
		}

		err := todoList.ToggleTodo(2)

		want := TodoList{
			{"Do laundry", true},
			{"Go shopping", false},
			{"learn go", true},
		}

		if !reflect.DeepEqual(todoList, want) {
			t.Errorf("Got %v want %v", todoList, want)
		}

		if err != nil {
			t.Errorf("Got an error %d", err)
		}
	})

	t.Run("Third item", func(t *testing.T) {
		todoList := TodoList{
			{"Do laundry", true},
			{"Go shopping", true},
			{"learn go", true},
		}

		err := todoList.ToggleTodo(3)

		want := TodoList{
			{"Do laundry", true},
			{"Go shopping", true},
			{"learn go", false},
		}

		if !reflect.DeepEqual(todoList, want) {
			t.Errorf("Got %v want %v", todoList, want)
		}

		if err != nil {
			t.Errorf("Got an error %d", err)
		}
	})

	t.Run("Index out of range", func(t *testing.T) {
		todoList := TodoList{
			{"Do laundry", false},
			{"Go shopping", false},
			{"learn go", false},
		}

		err := todoList.ToggleTodo(4)

		if err == nil {
			t.Errorf("Expected an error but didn't get one")
		}
	})
}
