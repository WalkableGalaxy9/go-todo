package gotodo

import (
	"bytes"
	"testing"
)

func TestAddItem(t *testing.T) {

	todoCLI := TodoViewModelCLI{}

	output := bytes.Buffer{}
	input := bytes.Buffer{}
	input.WriteString("New todo\n")

	todoCLI.TodoList = TodoList{}

	todoCLI.AddTodoInput(&input, &output)

	want := TodoList{
		{Title: "New todo", Complete: false},
	}

	AssertTodoList(want, todoCLI.TodoList, t)
	AssertString(output.String(), "Title:\n", t)
}

func TestPrintTodo(t *testing.T) {

	todoCLI := TodoViewModelCLI{}

	t.Run("Testing Incomplete Item", func(t *testing.T) {
		todoCLI.TodoList = TodoList{
			{"Do laundry", false},
			{"Go shopping", false},
			{"learn go", false},
		}

		buffer := bytes.Buffer{}
		todoCLI.PrintTodo(&buffer)

		want := "1. Do laundry - incomplete\n2. Go shopping - incomplete\n3. learn go - incomplete\n"

		AssertString(buffer.String(), want, t)
	})
	t.Run("Testing complete Item", func(t *testing.T) {
		todoCLI.TodoList = TodoList{
			{"Do laundry", true},
			{"Go shopping", true},
			{"learn go", true},
		}

		buffer := bytes.Buffer{}
		todoCLI.PrintTodo(&buffer)

		want := "1. Do laundry - complete\n2. Go shopping - complete\n3. learn go - complete\n"

		AssertString(buffer.String(), want, t)
	})
}

func TestDeleteItem(t *testing.T) {

	todoCLI := TodoViewModelCLI{}

	todoCLI.TodoList = TodoList{
		{"Do laundry", false},
		{"Go shopping", false},
		{"learn go", false},
	}

	output := bytes.Buffer{}
	input := bytes.Buffer{}
	input.WriteString("1\n")

	todoCLI.DeleteTodoInput(&input, &output)

	want := TodoList{
		{"Go shopping", false},
		{"learn go", false},
	}

	AssertTodoList(want, todoCLI.TodoList, t)
	AssertString(output.String(), "Number:\n", t)
}

func TestToggleItem(t *testing.T) {

	todoCLI := TodoViewModelCLI{}
	todoCLI.TodoList = TodoList{
		{"Do laundry", false},
		{"Go shopping", false},
		{"learn go", false},
	}

	output := bytes.Buffer{}
	input := bytes.Buffer{}
	input.WriteString("1\n")

	todoCLI.ToggleTodoInput(&input, &output)

	want := TodoList{
		{"Do laundry", true},
		{"Go shopping", false},
		{"learn go", false},
	}

	AssertTodoList(want, todoCLI.TodoList, t)
	AssertString(output.String(), "Number:\n", t)
}
