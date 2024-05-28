package gotodo

import (
	"bytes"
	"testing"
)

func TestPrintTodo(t *testing.T) {

	t.Run("Testing Incomplete Item", func(t *testing.T) {
		todo := TodoItem{
			Title:    "Do shopping",
			Complete: false,
		}

		buffer := bytes.Buffer{}
		PrintTodo(&buffer, todo)

		want := "Do shopping - incomplete"
		got := buffer.String()

		if got != want {
			t.Errorf("Want %s got %s", want, got)
		}
	})
	t.Run("Testing complete Item", func(t *testing.T) {
		todo := TodoItem{
			Title:    "Do shopping",
			Complete: true,
		}

		buffer := bytes.Buffer{}
		PrintTodo(&buffer, todo)

		want := "Do shopping - complete"
		got := buffer.String()

		if got != want {
			t.Errorf("Want %s got %s", want, got)
		}
	})
}
