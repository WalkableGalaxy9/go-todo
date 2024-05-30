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

		if AddCalls != 1 {
			t.Errorf("Expected a call to Add, got %d calls", AddCalls)
		}
	})

	t.Run("Delete item", func(t *testing.T) {
		resetCounts()
		input := bytes.Buffer{}
		output := bytes.Buffer{}
		input.WriteString("D\n")
		ManageTodoList(&spyViewModel, &input, &output)

		if DeleteCalls != 1 {
			t.Errorf("Expected a call to Delete, got %d calls", AddCalls)
		}
	})

	t.Run("Toggle item", func(t *testing.T) {
		resetCounts()
		input := bytes.Buffer{}
		output := bytes.Buffer{}
		input.WriteString("T\n")
		ManageTodoList(&spyViewModel, &input, &output)

		if ToggleCalls != 1 {
			t.Errorf("Expected a call to Toggle, got %d calls", AddCalls)
		}
	})

	t.Run("Garbage", func(t *testing.T) {
		resetCounts()
		input := bytes.Buffer{}
		output := bytes.Buffer{}
		input.WriteString("F\n")
		ManageTodoList(&spyViewModel, &input, &output)

		if PrintCalls != 1 || DeleteCalls != 0 || AddCalls != 0 || ToggleCalls != 0 {
			t.Errorf("Expected no calls, got %d Print calls, %d DeleteCalls, %d AddCalls, %d ToggleCalls ", PrintCalls, DeleteCalls, AddCalls, ToggleCalls)
		}
	})
}
