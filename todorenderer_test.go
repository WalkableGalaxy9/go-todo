package gotodo

import (
	"bytes"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {

	todoList := TodoList{
		{
			Title:    "Test",
			Complete: false,
		},
		{
			Title:    "Test2",
			Complete: true,
		},
		{
			Title:    "Test3",
			Complete: false,
		},
	}

	t.Run("Convert a single todo to Html", func(t *testing.T) {

		buffer := bytes.Buffer{}
		err := RenderTodo(&buffer, todoList)

		AssertNoError(err, t)

		got := buffer.String()
		approvals.VerifyString(t, got)
	})
}
