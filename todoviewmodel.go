package gotodo

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

func PrintTodo(writer io.Writer) {

	for index, item := range TodoList {
		completeString := "complete"

		if !item.Complete {
			completeString = "incomplete"
		}

		fmt.Fprintf(writer, "%d. %s - %s\n", index+1, item.Title, completeString)
	}
}

func AddTodoInput(input io.Reader, output io.Writer) {

	fmt.Fprintln(output, "Title:")

	reader := bufio.NewReader(input)
	title, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Error reading title: %v", err)
	}
	title = strings.TrimSpace(title)

	CreateTodo(title)
}

func DeleteTodoInput(input io.Reader, output io.Writer) {

	fmt.Fprintln(output, "Number:")

	reader := bufio.NewReader(input)
	title, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Error reading title: %v", err)
	}

	indexToRemove, _ := strconv.Atoi(strings.TrimSpace(title))

	DeleteTodo(indexToRemove)
}
