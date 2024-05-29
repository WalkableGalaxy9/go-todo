package gotodo

import (
	"reflect"
	"testing"
)

func TestCreateTodo(t *testing.T) {

	want := []TodoItem{
		{
			"Do laundry",
			false,
		},
		{
			"Learn Go",
			false,
		},
	}

	TodoList = []TodoItem{
		{
			"Do laundry",
			false,
		},
	}

	CreateTodo("Learn Go")

	if !reflect.DeepEqual(want, TodoList) {
		t.Errorf("Wanted %v got %v", want, TodoList)
	}

}

func TestDeleteTodo(t *testing.T) {

	t.Run("First item", func(t *testing.T) {
		TodoList = []TodoItem{
			{"Do laundry", false},
			{"Go shopping", false},
			{"learn go", false},
		}

		err := DeleteTodo(1)

		want := []TodoItem{
			{"Go shopping", false},
			{"learn go", false},
		}

		if !reflect.DeepEqual(TodoList, want) {
			t.Errorf("Got %v want %v", TodoList, want)
		}

		if err != nil {
			t.Errorf("Got an error %d", err)
		}
	})

	t.Run("Second item", func(t *testing.T) {
		TodoList = []TodoItem{
			{"Do laundry", false},
			{"Go shopping", false},
			{"learn go", false},
		}

		err := DeleteTodo(2)

		want := []TodoItem{
			{"Do laundry", false},
			{"learn go", false},
		}

		if !reflect.DeepEqual(TodoList, want) {
			t.Errorf("Got %v want %v", TodoList, want)
		}

		if err != nil {
			t.Errorf("Got an error %d", err)
		}
	})

	t.Run("Index out of range", func(t *testing.T) {
		TodoList = []TodoItem{
			{"Do laundry", false},
			{"Go shopping", false},
			{"learn go", false},
		}

		err := DeleteTodo(4)

		if err == nil {
			t.Errorf("Expected an error but didn't get one")
		}
	})
}
