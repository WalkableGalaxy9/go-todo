package gotodo

import (
	"bytes"
	"reflect"
	"testing"
)

func TestPrintTodo(t *testing.T) {

	t.Run("Testing Incomplete Item", func(t *testing.T) {
		todo := []TodoItem{
			{"Do laundry", false},
			{"Go shopping", false},
			{"learn go", false},
		}

		buffer := bytes.Buffer{}
		PrintTodo(&buffer, todo)

		want := "Do laundry - incomplete\nGo shopping - incomplete\nlearn go - incomplete\n"
		got := buffer.String()

		if got != want {
			t.Errorf("Want %s got %s", want, got)
		}
	})
	t.Run("Testing complete Item", func(t *testing.T) {
		todo := []TodoItem{
			{"Do laundry", true},
			{"Go shopping", true},
			{"learn go", true},
		}

		buffer := bytes.Buffer{}
		PrintTodo(&buffer, todo)

		want := "Do laundry - complete\nGo shopping - complete\nlearn go - complete\n"
		got := buffer.String()

		if got != want {
			t.Errorf("Want %s got %s", want, got)
		}
	})
}

func TestCreateList(t *testing.T) {

	want := []TodoItem{
		{"Do laundry", false},
		{"Go shopping", false},
		{"learn go", false},
	}

	got := CreateList("Do laundry", "Go shopping", "learn go")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Want %v got %v", want, got)
	}
}
