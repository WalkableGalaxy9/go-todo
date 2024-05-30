package gotodo

import (
	"bytes"
	"reflect"
	"testing"
)

func TestPrintTodo(t *testing.T) {

	t.Run("Testing Incomplete Item", func(t *testing.T) {
		TodoList = []TodoItem{
			{"Do laundry", false},
			{"Go shopping", false},
			{"learn go", false},
		}

		buffer := bytes.Buffer{}
		PrintTodo(&buffer)

		want := "1. Do laundry - incomplete\n2. Go shopping - incomplete\n3. learn go - incomplete\n"
		got := buffer.String()

		if got != want {
			t.Errorf("Want %s got %s", want, got)
		}
	})
	t.Run("Testing complete Item", func(t *testing.T) {
		TodoList = []TodoItem{
			{"Do laundry", true},
			{"Go shopping", true},
			{"learn go", true},
		}

		buffer := bytes.Buffer{}
		PrintTodo(&buffer)

		want := "1. Do laundry - complete\n2. Go shopping - complete\n3. learn go - complete\n"
		got := buffer.String()

		if got != want {
			t.Errorf("Want %s got %s", want, got)
		}
	})
}

func TestAddItem(t *testing.T) {

	output := bytes.Buffer{}
	input := bytes.Buffer{}
	input.WriteString("New todo")

	AddTodoInput(&input, &output)

	want := []TodoItem{{Title: "New todo", Complete: false}}

	if !reflect.DeepEqual(TodoList, want) {
		t.Errorf("Got %v want %v", TodoList, want)
	}

	if output.String() != "Title:\n" {
		t.Errorf("Got %v want %v", output.String(), "Title:\n")
	}
}

func TestDeleteItem(t *testing.T) {

	TodoList = []TodoItem{
		{"Do laundry", false},
		{"Go shopping", false},
		{"learn go", false},
	}

	output := bytes.Buffer{}
	input := bytes.Buffer{}
	input.WriteString("1\n")

	DeleteTodoInput(&input, &output)

	want := []TodoItem{
		{"Go shopping", false},
		{"learn go", false},
	}

	if !reflect.DeepEqual(TodoList, want) {
		t.Errorf("Got %v want %v", TodoList, want)
	}

	if output.String() != "Number:\n" {
		t.Errorf("Got %v want %v", output.String(), "Number:\n")
	}
}

func TestToggleItem(t *testing.T) {

	TodoList = []TodoItem{
		{"Do laundry", false},
		{"Go shopping", false},
		{"learn go", false},
	}

	output := bytes.Buffer{}
	input := bytes.Buffer{}
	input.WriteString("1\n")

	ToggleTodoInput(&input, &output)

	want := []TodoItem{
		{"Do laundry", true},
		{"Go shopping", false},
		{"learn go", false},
	}

	if !reflect.DeepEqual(TodoList, want) {
		t.Errorf("Got %v want %v", TodoList, want)
	}

	if output.String() != "Number:\n" {
		t.Errorf("Got %v want %v", output.String(), "Number:\n")
	}
}
