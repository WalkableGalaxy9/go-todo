package gotodo

import (
	"bytes"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {

	todoHTML := TodoRenderHTML{}

	todoHTML.TodoList = TodoList{
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
		err := todoHTML.RenderTodo(&buffer)

		AssertNoError(err, t)

		got := buffer.String()
		approvals.VerifyString(t, got)
	})
}

func TestRenderCreate(t *testing.T) {

	todoHTML := TodoRenderHTML{}

	todoHTML.TodoList = TodoList{
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

	t.Run("Render the new Todo Page", func(t *testing.T) {

		buffer := bytes.Buffer{}
		err := todoHTML.RenderCreate(&buffer)

		AssertNoError(err, t)

		got := buffer.String()
		approvals.VerifyString(t, got)
	})
}

func TestCreateTodoFromPage(t *testing.T) {

	todoHTML := TodoRenderHTML{}

	todoHTML.TodoList = TodoList{
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
	want := TodoList{
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
		{
			Title:    "Test4",
			Complete: false,
		},
	}

	t.Run("Adding a new todo", func(t *testing.T) {
		err := todoHTML.CreateTodoFromPage("Test4")

		AssertNoError(err, t)

		AssertTodoList(want, todoHTML.TodoList, t)
	})
}
