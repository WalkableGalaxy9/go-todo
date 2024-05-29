package gotodo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"
)

type MenuOption int

const (
	MenuUnknown = iota
	MenuAddTodo
)

type TodoItem struct {
	Title    string
	Complete bool
}

var TodoList []TodoItem

func PrintTodo(writer io.Writer) {

	for index, item := range TodoList {
		completeString := "complete"

		if !item.Complete {
			completeString = "incomplete"
		}

		fmt.Fprintf(writer, "%d. %s - %s\n", index+1, item.Title, completeString)
	}
}

func CreateList(items ...string) {

	list := []TodoItem{}
	for _, item := range items {
		list = append(list, TodoItem{item, false})
	}

	TodoList = list

}

func PrintTodoJSON(writer io.Writer) {

	json.NewEncoder(writer).Encode(TodoList)
}

func WriteJSONToFile(filename string) error {

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return err
	}
	defer file.Close()

	PrintTodoJSON(file)

	return nil
}

func ReadJSONFromAFile(filesystem fs.FS, filename string, writer io.Writer) error {
	err := ExtractItemsFromJSONFile(filesystem, filename)
	if err != nil {
		return err
	}

	PrintTodo(writer)

	return nil
}

func ExtractItemsFromJSONFile(filesystem fs.FS, filename string) error {
	todofile, err := filesystem.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return err
	}
	defer todofile.Close()

	data, err := io.ReadAll(todofile)

	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return err
	}

	err = json.Unmarshal(data, &TodoList)

	if err != nil {
		fmt.Printf("Error parsing JSON: %v", err)
		return err
	}

	return nil
}

func CreateTodo(title string, complete bool) TodoItem {
	return TodoItem{Title: title, Complete: complete}
}

func AddTodoInput(input io.Reader, output io.Writer) {

	fmt.Fprintln(output, "Title:")

	reader := bufio.NewReader(input)
	title, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Error reading title: %v", err)
	}
	title = strings.TrimSpace(title)

	TodoList = append(TodoList, CreateTodo(title, false))
}

func GetMenuOption(input io.Reader, output io.Writer) MenuOption {

	fmt.Fprintf(output, "A. Add Todo\n")

	reader := bufio.NewReader(input)
	option, _, err := reader.ReadRune()

	if err != nil {
		log.Fatalf("Error reading title: %v", err)
	}

	switch option {
	case 'A':
		return MenuAddTodo
	default:
		return MenuUnknown
	}
}
