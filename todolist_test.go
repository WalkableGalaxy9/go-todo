package gotodo

import (
	"testing"
)

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

	AssertTodoList(want, todoList, t)
}

func TestDeleteTodo(t *testing.T) {

	t.Run("Delete first item", func(t *testing.T) {
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

		AssertTodoList(want, todoList, t)
		AssertNoError(err, t)
	})

	t.Run("Delete second item", func(t *testing.T) {
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

		AssertTodoList(want, todoList, t)
		AssertNoError(err, t)
	})

	t.Run("Delete Third item", func(t *testing.T) {
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

		AssertTodoList(want, todoList, t)
		AssertNoError(err, t)
	})

	t.Run("Delete index out of range", func(t *testing.T) {
		todoList := TodoList{
			{"Do laundry", false},
			{"Go shopping", false},
			{"learn go", false},
		}

		err := todoList.DeleteTodo(4)

		AssertError(err, t)
	})
}

func TestToggleTodo(t *testing.T) {

	t.Run("Toggle first item", func(t *testing.T) {
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

		AssertTodoList(want, todoList, t)
		AssertNoError(err, t)
	})

	t.Run("Toggle Second item", func(t *testing.T) {
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

		AssertTodoList(want, todoList, t)
		AssertNoError(err, t)
	})

	t.Run("Toggle third item", func(t *testing.T) {
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

		AssertTodoList(want, todoList, t)
		AssertNoError(err, t)
	})

	t.Run("Toggle index out of range", func(t *testing.T) {
		todoList := TodoList{
			{"Do laundry", false},
			{"Go shopping", false},
			{"learn go", false},
		}

		err := todoList.ToggleTodo(4)
		AssertError(err, t)
	})
}
