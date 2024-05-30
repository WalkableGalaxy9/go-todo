package gotodo

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

func PrintTodo(list *TodoList, writer io.Writer) {

	for index, item := range *list {
		completeString := "complete"

		if !item.Complete {
			completeString = "incomplete"
		}

		fmt.Fprintf(writer, "%d. %s - %s\n", index+1, item.Title, completeString)
	}
}

func AddTodoInput(list *TodoList, input io.Reader, output io.Writer) {

	fmt.Fprintln(output, "Title:")

	reader := bufio.NewReader(input)
	title, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Error reading title: %v", err)
	}
	title = strings.TrimSpace(title)

	list.CreateTodo(title)
}

func DeleteTodoInput(list *TodoList, input io.Reader, output io.Writer) {

	fmt.Fprintln(output, "Number:")

	reader := bufio.NewReader(input)
	title, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Error reading title: %v", err)
	}

	indexToRemove, _ := strconv.Atoi(strings.TrimSpace(title))

	list.DeleteTodo(indexToRemove)
}

func ToggleTodoInput(list *TodoList, input io.Reader, output io.Writer) {

	fmt.Fprintln(output, "Number:")

	reader := bufio.NewReader(input)
	title, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Error reading title: %v", err)
	}

	index, _ := strconv.Atoi(strings.TrimSpace(title))

	list.ToggleTodo(index)
}
