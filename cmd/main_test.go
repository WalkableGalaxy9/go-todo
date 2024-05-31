package main

import (
	"bytes"
	"io"
	"testing"
)

type SpyTodoViewModel struct {
}

var PrintCalls = 0

func (s *SpyTodoViewModel) PrintTodo(io.Writer) {
	PrintCalls++
}

var AddCalls = 0

func (s *SpyTodoViewModel) AddTodoInput(io.Reader, io.Writer) {
	AddCalls++
}

var DeleteCalls = 0

func (s *SpyTodoViewModel) DeleteTodoInput(io.Reader, io.Writer) {
	DeleteCalls++
}

var ToggleCalls = 0

func (s *SpyTodoViewModel) ToggleTodoInput(io.Reader, io.Writer) {
	ToggleCalls++
}

func resetCounts() {
	PrintCalls = 0
	AddCalls = 0
	DeleteCalls = 0
	ToggleCalls = 0
}

func TestManageTodoList(t *testing.T) {

	spyViewModel := SpyTodoViewModel{}

	t.Run("Add item", func(t *testing.T) {
		resetCounts()

		input := bytes.Buffer{}
		output := bytes.Buffer{}

		input.WriteString("A\n")
		ManageTodoList(&spyViewModel, &input, &output)

		assertCalls(t, 1, 0, 0, 1)
	})

	t.Run("Delete item", func(t *testing.T) {
		resetCounts()
		input := bytes.Buffer{}
		output := bytes.Buffer{}
		input.WriteString("D\n")
		ManageTodoList(&spyViewModel, &input, &output)

		assertCalls(t, 0, 1, 0, 1)
	})

	t.Run("Toggle item", func(t *testing.T) {
		resetCounts()
		input := bytes.Buffer{}
		output := bytes.Buffer{}
		input.WriteString("T\n")
		ManageTodoList(&spyViewModel, &input, &output)

		assertCalls(t, 0, 0, 1, 1)
	})

	t.Run("Garbage", func(t *testing.T) {
		resetCounts()
		input := bytes.Buffer{}
		output := bytes.Buffer{}
		input.WriteString("F\n")
		ManageTodoList(&spyViewModel, &input, &output)

		assertCalls(t, 0, 0, 0, 1)
	})
}

func assertCalls(t *testing.T, wantAdd, wantDelete, wantToggle, wantPrint int) {
	t.Helper()
	assertCall(t, wantAdd, AddCalls, "Add")
	assertCall(t, wantDelete, DeleteCalls, "Delete")
	assertCall(t, wantToggle, ToggleCalls, "Toggle")
	assertCall(t, wantPrint, PrintCalls, "Print")
}

func assertCall(t *testing.T, want, got int, name string) {
	t.Helper()
	if want != got {
		t.Errorf("Expected %d calls to %s, got %d calls", want, name, got)
	}
}
