package gotodo

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type TodoViewModel interface {
	PrintTodo(io.Writer)
	AddTodoInput(io.Reader, io.Writer)
	DeleteTodoInput(io.Reader, io.Writer)
	ToggleTodoInput(io.Reader, io.Writer)
}

type TodoViewModelCLI struct {
	TodoList TodoList
}

func (t *TodoViewModelCLI) PrintTodo(writer io.Writer) {

	for index, item := range t.TodoList {
		completeString := "complete"

		if !item.Complete {
			completeString = "incomplete"
		}

		fmt.Fprintf(writer, "%d. %s - %s\n", index+1, item.Title, completeString)
	}
}

func (t *TodoViewModelCLI) AddTodoInput(input io.Reader, output io.Writer) {

	fmt.Fprintln(output, "Title:")

	reader := bufio.NewReader(input)
	title, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Error reading title: %v", err)
	}
	title = strings.TrimSpace(title)

	t.TodoList.CreateTodo(title)
}

func (t *TodoViewModelCLI) DeleteTodoInput(input io.Reader, output io.Writer) {

	fmt.Fprintln(output, "Number:")

	reader := bufio.NewReader(input)
	title, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Error reading title: %v", err)
	}

	indexToRemove, _ := strconv.Atoi(strings.TrimSpace(title))

	t.TodoList.DeleteTodo(indexToRemove)
}

func (t *TodoViewModelCLI) ToggleTodoInput(input io.Reader, output io.Writer) {

	fmt.Fprintln(output, "Number:")

	reader := bufio.NewReader(input)
	title, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Error reading title: %v", err)
	}

	index, _ := strconv.Atoi(strings.TrimSpace(title))

	t.TodoList.ToggleTodo(index)
}
