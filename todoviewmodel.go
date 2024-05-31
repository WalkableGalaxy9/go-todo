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

const (
	completeString   = "complete"
	incompleteString = "incomplete"
)

func (t *TodoViewModelCLI) PrintTodo(output io.Writer) {

	for index, todo := range t.TodoList {
		statusString := completeString

		if !todo.Complete {
			statusString = incompleteString
		}

		fmt.Fprintf(output, "%d. %s - %s\n", index+1, todo.Title, statusString)
	}
}

func (t *TodoViewModelCLI) AddTodoInput(input io.Reader, output io.Writer) {

	fmt.Fprintln(output, "Title:")

	title := getInput(input)

	t.TodoList.CreateTodo(title)
}

func (t *TodoViewModelCLI) DeleteTodoInput(input io.Reader, output io.Writer) {

	fmt.Fprintln(output, "Number:")

	title := getInput(input)

	indexToRemove, _ := strconv.Atoi(title)

	t.TodoList.DeleteTodo(indexToRemove)
}

func (t *TodoViewModelCLI) ToggleTodoInput(input io.Reader, output io.Writer) {

	fmt.Fprintln(output, "Number:")

	title := getInput(input)

	index, _ := strconv.Atoi(title)

	t.TodoList.ToggleTodo(index)
}

func getInput(input io.Reader) string {
	reader := bufio.NewReader(input)
	title, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	title = strings.TrimSpace(title)

	return title
}
